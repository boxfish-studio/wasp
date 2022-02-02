package vmcontext

import (
	"github.com/iotaledger/wasp/packages/iscp/coreutil"
	"github.com/iotaledger/wasp/packages/vm/gas"
	"github.com/iotaledger/wasp/packages/vm/vmcontext/vmtxbuilder"
	"golang.org/x/xerrors"
)

func (vmctx *VMContext) gasBurnEnable(enable bool) {
	vmctx.gasBurnEnabled = enable
}

func (vmctx *VMContext) gasSetBudget(gasBudget uint64) {
	vmctx.gasBudgetAdjusted = gasBudget
	vmctx.gasBurned = 0
}

func (vmctx *VMContext) GasBurn(burnCode gas.BurnCode, par ...uint64) {
	if !vmctx.gasBurnEnabled {
		return
	}
	g := burnCode.Cost(par...)
	vmctx.gasBurnLog.Record(burnCode, g)
	vmctx.gasBurned += g

	if vmctx.gasBurnedTotal+g > gas.MaxGasPerBlock {
		panic(vmtxbuilder.ErrBlockGasLimitExceeded)
	}
	vmctx.gasBurnedTotal += g

	if vmctx.gasBurned > vmctx.gasBudgetAdjusted {
		panic(xerrors.Errorf("%v: burned (budget) = %d (%d)",
			coreutil.ErrorGasBudgetExceeded, vmctx.gasBurned, vmctx.gasBudgetAdjusted))
	}
}

func (vmctx *VMContext) GasBudgetLeft() uint64 {
	if vmctx.gasBudgetAdjusted < vmctx.gasBurned {
		return 0
	}
	return vmctx.gasBudgetAdjusted - vmctx.gasBurned
}

func (vmctx *VMContext) GasBurned() uint64 {
	return vmctx.gasBurned
}