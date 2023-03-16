<script lang="ts">
  import { toast } from '@zerodevx/svelte-toast';

  import { Button, Input } from '$components';

  import { Bech32AddressLength, EVMAddressLength } from '$lib/constants';
  import { IotaWallet, SendFundsTransaction } from '$lib/faucet';
  import { indexerClient, nodeClient, selectedNetwork } from '$lib/../store';

  let isSendingFunds: boolean;

  let balance: bigint = BigInt(0);
  let evmAddress: string = '';

  $: enableSendFunds =
    evmAddress.length == EVMAddressLength &&
    $selectedNetwork != null &&
    $selectedNetwork.chainAddress.length == Bech32AddressLength &&
    !isSendingFunds;

  async function sendFunds() {
    if (!enableSendFunds) {
      return;
    }

    isSendingFunds = true;

    let wallet: IotaWallet = new IotaWallet(
      $nodeClient,
      $indexerClient,
      $selectedNetwork.faucetEndpoint,
    );

    let toastId: number;

    try {
      toastId = toast.push('Initializing wallet');
      await wallet.initialize();
      toast.pop(toastId);

      toastId = toast.push('Requesting funds from the faucet', {
        duration: 20 * 2000, // 20 retries, 2s delay each.
      });
      balance = await wallet.requestFunds();
      toast.pop(toastId);

      toastId = toast.push('Sending funds');
      const transaction = new SendFundsTransaction(wallet);
      await transaction.sendFundsToEVMAddress(
        evmAddress,
        $selectedNetwork.chainAddress,
        balance,
        BigInt(5000000),
      );
      toast.pop(toastId);

      toast.push(
        'Funds successfully sent! It may take 10-30 seconds to arive.',
        {
          duration: 10 * 1000,
        },
      );
    } catch (ex) {
      toast.pop(toastId);
      toast.push(ex.message);
    }

    isSendingFunds = false;
  }
</script>

<faucet-component class="flex flex-col space-y-6 mt-6">
  {#if $selectedNetwork}
    <Input
      id="evmAddress"
      label="EVM Address"
      bind:value={evmAddress}
      placeholder="0x..."
      stretch
      autofocus
    />

    <Button
      title="Send funds"
      disabled={!enableSendFunds}
      onClick={sendFunds}
      busy={isSendingFunds}
    />
  {:else}
    <span> Please select a network first. </span>
  {/if}
</faucet-component>

<style>
  .error {
    background-color: #9e534a47;
    border: 2px solid #991c0d78;
    border-radius: 10px;
    padding: 15px;
  }

  .error_title {
    font-weight: bold;
    margin-bottom: 15px;
  }

  component {
    color: rgba(255, 255, 255, 0.87);
    display: flex;
    flex-direction: column;
  }
</style>
