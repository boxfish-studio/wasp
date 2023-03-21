import type { IndexerPluginClient, SingleNodeClient } from '@iota/iota.js';
import { Converter } from '@iota/util.js';
import {
  Multicall, type ContractCallContext, type ContractCallResults
} from 'ethereum-multicall';
import type Web3 from 'web3';
import type { Eth } from 'web3-eth';
import type { Contract } from 'web3-eth-contract';

import { NativeTokenIDLength } from '$lib/constants';
import { evmAddressToAgentID } from '$lib/evm';
import { hNameFromString } from '$lib/hname';
import { getNativeTokenMetaData, type INativeToken } from '$lib/native_token';
import { getNFTMetadata, type INFT } from '$lib/nft';
import { gasFee, getBalanceParameters, iscAbi, iscContractAddress, withdrawParameters } from '$lib/withdraw';

export type NFTDict = [string, string][];

export interface IWithdrawResponse {
  blockHash: string;
  blockNumber: number;
  contractAddress: string;
  cumulativeGasUsed: number;
  from: string;
  gasUsed: number;
  logsBloom: string;
  status: boolean;
  to: string;
  transactionHash: string;
  transactionIndex: number;
  events: unknown;
}

export class ISCMagic {
  private readonly contract: Contract;
  private readonly multicall: Contract;

  constructor(contract: Contract, multicall: Contract) {
    this.contract = contract;
    this.multicall = multicall;
  }

  public async getBaseTokens(eth: Eth, account: string): Promise<number> {
    const addressBalance = await eth.getBalance(account);
    const balance = BigInt(addressBalance) / BigInt(1e12);

    return Number(balance);
  }

  public async getNativeTokens(nodeClient: SingleNodeClient, indexerClient: IndexerPluginClient, account: string) {
    const accountsCoreContract = hNameFromString('accounts');
    const getBalanceFunc = hNameFromString('balance');
    const agentID = evmAddressToAgentID(account);

    const parameters = getBalanceParameters(agentID);

    const nativeTokenResult = await this.contract.methods
      .callView(accountsCoreContract, getBalanceFunc, parameters)
      .call();

    const nativeTokens: INativeToken[] = [];

    for (let item of nativeTokenResult.items) {
      const id = item.key;
      const idBytes = Converter.hexToBytes(id);

      if (idBytes.length != NativeTokenIDLength) {
        continue;
      }

      var nativeToken: INativeToken = {
        // TODO: BigInt is required for native tokens, but it causes problems with the range slider. This needs to be adressed before shipping.
        amount: BigInt(item.value),
        id: id,
        metadata: await getNativeTokenMetaData(nodeClient, indexerClient, id),
      };

      nativeTokens.push(nativeToken);
    }

    return nativeTokens;
  }

  public async getNFTs(nodeClient: SingleNodeClient, indexerClient: IndexerPluginClient, account: string) {
    const accountsCoreContract = hNameFromString('accounts');
    const getAccountNFTsFunc = hNameFromString('accountNFTs');
    const agentID = evmAddressToAgentID(account);

    let parameters = getBalanceParameters(agentID);

    const NFTsResult = await this.contract.methods
      .callView(accountsCoreContract, getAccountNFTsFunc, parameters)
      .call();

    const nfts = NFTsResult.items as NFTDict;

    // The 'i' parameter returns the length of the nft id array, but we can just filter that out
    // and go through the list dynamically.
    const nftIds = nfts.filter(x => Converter.hexToUtf8(x[0]) != 'i');
    const availableNFTs = nftIds.map(x => <INFT>{ id: x[1] });

    console.log(nftIds)

    for (let nft of availableNFTs) {
      nft.metadata = await getNFTMetadata(nodeClient, indexerClient, nft.id);
    }

    return availableNFTs;
  }

  public async withdrawMulticall(web3Instance: Web3, multicallAddress: string, nodeClient: SingleNodeClient, receiverAddress: string, baseTokens: number, nativeTokens: INativeToken[], nftIDs: string[]) {
    const sendABI = iscAbi.find(x => x.name == 'send');
    const contractCallContext: ContractCallContext =
    {
      reference: 'send',
      contractAddress: iscContractAddress,

      abi: iscAbi,
      calls: []
    };

    let baseTokensSent = 0;

    for (let nft of nftIDs) {
      const parameters = await withdrawParameters(
        nodeClient,
        receiverAddress,
        gasFee,
        gasFee * 2,
        [],
        nft,
      );

      baseTokensSent += gasFee * 2;

      contractCallContext.calls.push({
        reference: 'send', methodName: 'send', methodParameters: parameters,
      })
    }

    const lastParameters = await withdrawParameters(
      nodeClient,
      receiverAddress,
      gasFee,
      baseTokens - baseTokensSent,
      nativeTokens,
      null,
    );

    console.log(contractCallContext)

    /*contractCallContext.calls.push({
      reference: 'send', methodName: 'send', methodParameters: lastParameters,
    });*/

    const multicall = new Multicall({ web3Instance: web3Instance, tryAggregate: false, multicallCustomContractAddress: multicallAddress });
    const results: ContractCallResults = await multicall.call(contractCallContext);

    console.log(results);

  }

  public async withdraw(nodeClient: SingleNodeClient, receiverAddress: string, baseTokens: number, nativeTokens: INativeToken[], nftID?: string) {
    const parameters = await withdrawParameters(
      nodeClient,
      receiverAddress,
      gasFee,
      baseTokens,
      nativeTokens,
      nftID,
    );

    let result = await this.contract.methods.send(...parameters).send();

    return result as IWithdrawResponse;
  }
}