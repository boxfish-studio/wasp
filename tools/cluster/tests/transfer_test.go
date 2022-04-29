package tests

import (
	"testing"
	"time"

	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/utxodb"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/stretchr/testify/require"
)

func TestDepositWithdraw(t *testing.T) {
	e := setupWithNoChain(t)

	chain, err := e.clu.DeployDefaultChain()
	require.NoError(t, err)

	chEnv := newChainEnv(t, e.clu, chain)

	myWallet, myAddress, err := e.clu.NewKeyPairWithFunds()
	require.NoError(e.t, err)

	require.True(t,
		e.clu.AssertAddressBalances(myAddress, iscp.NewTokensIotas(utxodb.FundsFromFaucetAmount)),
	)
	chEnv.checkLedger()

	myAgentID := iscp.NewAgentID(myAddress, 0)
	// origAgentID := iscp.NewAgentID(chain.OriginatorAddress(), 0)

	// chEnv.checkBalanceOnChain(origAgentID, iscp.IotaTokenID, 0)
	chEnv.checkBalanceOnChain(myAgentID, iscp.IotaTokenID, 0)
	chEnv.checkLedger()

	// deposit some iotas to the chain
	depositIotas := uint64(42000)
	chClient := chainclient.New(e.clu.L1Client(), e.clu.WaspClient(0), chain.ChainID, myWallet)

	par := chainclient.NewPostRequestParams().WithIotas(depositIotas)
	reqTx, err := chClient.Post1Request(accounts.Contract.Hname(), accounts.FuncDeposit.Hname(), *par)
	require.NoError(t, err)

	receipts, err := chain.CommitteeMultiClient().WaitUntilAllRequestsProcessedSuccessfully(chain.ChainID, reqTx, 30*time.Second)
	require.NoError(t, err)
	chEnv.checkLedger()

	// chEnv.checkBalanceOnChain(origAgentID, iscp.IotaTokenID, 0)
	gasFees1 := receipts[0].GasFeeCharged
	onChainBalance := depositIotas - gasFees1
	chEnv.checkBalanceOnChain(myAgentID, iscp.IotaTokenID, onChainBalance)

	require.True(t,
		e.clu.AssertAddressBalances(myAddress, iscp.NewTokensIotas(utxodb.FundsFromFaucetAmount-depositIotas)),
	)

	// withdraw some iotas back
	iotasToWithdraw := uint64(500)
	req, err := chClient.PostOffLedgerRequest(accounts.Contract.Hname(), accounts.FuncWithdraw.Hname(),
		chainclient.PostRequestParams{
			Allowance: iscp.NewAllowanceIotas(iotasToWithdraw),
		},
	)
	require.NoError(t, err)
	receipt, err := chain.CommitteeMultiClient().WaitUntilRequestProcessedSuccessfully(chain.ChainID, req.ID(), 30*time.Second)
	require.NoError(t, err)

	chEnv.checkLedger()
	gasFees2 := receipt.GasFeeCharged
	chEnv.checkBalanceOnChain(myAgentID, iscp.IotaTokenID, onChainBalance-iotasToWithdraw-gasFees2)
	require.True(t,
		e.clu.AssertAddressBalances(myAddress, iscp.NewTokensIotas(utxodb.FundsFromFaucetAmount-depositIotas+iotasToWithdraw)),
	)

	// TODO use "withdraw all base tokens" entrypoint to withdraw all remaining iotas
	t.Fatal()
}
