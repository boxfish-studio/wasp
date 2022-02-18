package wasmsolo

import (
	"bytes"
	"errors"
	"time"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/gas"
	"github.com/iotaledger/wasp/packages/vm/sandbox"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmhost"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmrequests"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
	"golang.org/x/xerrors"
)

// NOTE: These functions correspond to the Sandbox fnXxx constants in WasmLib
var sandboxFunctions = []func(*SoloSandbox, []byte) []byte{
	nil,
	(*SoloSandbox).fnAccountID,
	(*SoloSandbox).fnAllowance,
	(*SoloSandbox).fnBalance,
	(*SoloSandbox).fnBalances,
	(*SoloSandbox).fnBlockContext,
	(*SoloSandbox).fnCall,
	(*SoloSandbox).fnCaller,
	(*SoloSandbox).fnChainID,
	(*SoloSandbox).fnChainOwnerID,
	(*SoloSandbox).fnContract,
	(*SoloSandbox).fnContractCreator,
	(*SoloSandbox).fnDeployContract,
	(*SoloSandbox).fnEntropy,
	(*SoloSandbox).fnEvent,
	(*SoloSandbox).fnLog,
	(*SoloSandbox).fnMinted,
	(*SoloSandbox).fnPanic,
	(*SoloSandbox).fnParams,
	(*SoloSandbox).fnPost,
	(*SoloSandbox).fnRequest,
	(*SoloSandbox).fnRequestID,
	(*SoloSandbox).fnResults,
	(*SoloSandbox).fnSend,
	(*SoloSandbox).fnStateAnchor,
	(*SoloSandbox).fnTimestamp,
	(*SoloSandbox).fnTrace,
	(*SoloSandbox).fnUtilsBase58Decode,
	(*SoloSandbox).fnUtilsBase58Encode,
	(*SoloSandbox).fnUtilsBlsAddress,
	(*SoloSandbox).fnUtilsBlsAggregate,
	(*SoloSandbox).fnUtilsBlsValid,
	(*SoloSandbox).fnUtilsEd25519Address,
	(*SoloSandbox).fnUtilsEd25519Valid,
	(*SoloSandbox).fnUtilsHashBlake2b,
	(*SoloSandbox).fnUtilsHashName,
	(*SoloSandbox).fnUtilsHashSha3,
	(*SoloSandbox).fnTransferAllowed,
}

// SoloSandbox acts as a temporary host side of the WasmLib Sandbox interface.
// It acts as a change-resistant layer to wrap changes to the Solo environment,
// to limit bothering users of WasmLib as little as possible with those changes.
// Note that only those functions that are related to invocation of SC requests
// are actually necessary here. These sandbox functions will never be called
// other than through the SC function call interface generated by schema tool.
type SoloSandbox struct {
	ctx   *SoloContext
	cvt   wasmhost.WasmConvertor
	utils iscp.Utils
}

func (s *SoloSandbox) Burn(burnCode gas.BurnCode, par ...uint64) {
	panic("implement me")
}

func (s *SoloSandbox) Budget() uint64 {
	panic("implement me")
}

var (
	_ wasmhost.ISandbox = new(SoloSandbox)
	_ iscp.Gas          = new(SoloSandbox)
)

func NewSoloSandbox(ctx *SoloContext) *SoloSandbox {
	s := &SoloSandbox{ctx: ctx}
	s.utils = sandbox.NewUtils(s)
	return s
}

func (s *SoloSandbox) Call(funcNr int32, params []byte) []byte {
	s.ctx.Err = nil
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		if s.ctx.Err != nil {
			s.ctx.Chain.Log().Infof("stacked error: %s", s.ctx.Err.Error())
		}
		switch errType := r.(type) {
		case error:
			s.ctx.Err = errType
		case string:
			s.ctx.Err = errors.New(errType)
		default:
			s.ctx.Err = xerrors.Errorf("RunScFunction: %v", errType)
		}
		s.ctx.Chain.Log().Infof("stolor error:: %s", s.ctx.Err.Error())
	}()
	return sandboxFunctions[-funcNr](s, params)
}

func (s *SoloSandbox) checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (s *SoloSandbox) Panicf(format string, args ...interface{}) {
	s.ctx.Chain.Log().Panicf(format, args...)
}

func (s *SoloSandbox) Tracef(format string, args ...interface{}) {
	s.ctx.Chain.Log().Debugf(format, args...)
}

func (s *SoloSandbox) postSync(contract, function string, params dict.Dict, assets *iscp.Assets) []byte {
	req := solo.NewCallParamsFromDic(contract, function, params)
	req.AddAssets(assets)
	req.WithAllowance(assets)
	req.WithMaxAffordableGasBudget()
	ctx := s.ctx
	//TODO mint!
	//if ctx.mint > 0 {
	//	mintAddress := ledgerstate.NewED25519Address(ctx.keyPair.PublicKey)
	//	req.WithMint(mintAddress, ctx.mint)
	//}
	_ = wasmhost.Connect(ctx.wasmHostOld)
	var res dict.Dict
	if ctx.offLedger {
		ctx.offLedger = false
		res, ctx.Err = ctx.Chain.PostRequestOffLedger(req, ctx.keyPair)
	} else if !ctx.isRequest {
		ctx.Tx, res, ctx.Err = ctx.Chain.PostRequestSyncTx(req, ctx.keyPair)
	} else {
		ctx.isRequest = false
		ctx.Tx, _, ctx.Err = ctx.Chain.RequestFromParamsToLedger(req, nil)
		if ctx.Err == nil {
			ctx.Chain.Env.EnqueueRequests(ctx.Tx)
		}
	}
	_ = wasmhost.Connect(ctx.wc)
	ctx.UpdateGas()
	if ctx.Err != nil {
		return nil
	}
	return res.Bytes()
}

//////////////////// sandbox functions \\\\\\\\\\\\\\\\\\\\

func (s *SoloSandbox) fnAccountID(args []byte) []byte {
	return s.ctx.AccountID().Bytes()
}

func (s *SoloSandbox) fnAllowance(args []byte) []byte {
	//// zero incoming balance
	assets := new(iscp.Assets)
	return s.cvt.ScBalances(assets).Bytes()
}

func (s *SoloSandbox) fnBalance(args []byte) []byte {
	color := wasmtypes.ColorFromBytes(args)
	return codec.EncodeUint64(s.ctx.Balance(s.ctx.Account(), color))
}

func (s *SoloSandbox) fnBalances(args []byte) []byte {
	agent := s.ctx.Account()
	account := iscp.NewAgentID(agent.address, agent.hname)
	assets := s.ctx.Chain.L2Assets(account)
	return s.cvt.ScBalances(assets).Bytes()
}

func (s *SoloSandbox) fnBlockContext(args []byte) []byte {
	panic("implement me")
}

func (s *SoloSandbox) fnCall(args []byte) []byte {
	ctx := s.ctx
	req := wasmrequests.NewCallRequestFromBytes(args)
	contract := s.cvt.IscpHname(req.Contract)
	if contract != iscp.Hn(ctx.scName) {
		s.Panicf("unknown contract: %s vs. %s", contract.String(), ctx.scName)
	}
	function := s.cvt.IscpHname(req.Function)
	funcName := ctx.wc.FunctionFromCode(uint32(function))
	if funcName == "" {
		s.Panicf("unknown function: %s", function.String())
	}
	s.Tracef("CALL %s.%s", ctx.scName, funcName)
	params, err := dict.FromBytes(req.Params)
	s.checkErr(err)
	scAssets := wasmlib.NewScAssetsFromBytes(req.Transfer)
	if len(scAssets) != 0 {
		assets := s.cvt.IscpAssets(scAssets)
		return s.postSync(ctx.scName, funcName, params, assets)
	}

	_ = wasmhost.Connect(ctx.wasmHostOld)
	res, err := ctx.Chain.CallView(ctx.scName, funcName, params)
	_ = wasmhost.Connect(ctx.wc)
	ctx.Err = err
	ctx.UpdateGas()
	if ctx.Err != nil {
		return nil
	}
	return res.Bytes()
}

func (s *SoloSandbox) fnCaller(args []byte) []byte {
	return s.ctx.Chain.OriginatorAgentID.Bytes()
}

func (s *SoloSandbox) fnChainID(args []byte) []byte {
	return s.ctx.ChainID().Bytes()
}

func (s *SoloSandbox) fnChainOwnerID(args []byte) []byte {
	return s.ctx.ChainOwnerID().Bytes()
}

func (s *SoloSandbox) fnContract(args []byte) []byte {
	return s.ctx.Account().hname.Bytes()
}

func (s *SoloSandbox) fnContractCreator(args []byte) []byte {
	return s.ctx.ContractCreator().Bytes()
}

func (s *SoloSandbox) fnDeployContract(args []byte) []byte {
	panic("implement me")
}

func (s *SoloSandbox) fnEntropy(args []byte) []byte {
	return s.ctx.Chain.ChainID.Bytes()
}

func (s *SoloSandbox) fnEvent(args []byte) []byte {
	s.Panicf("solo cannot send events")
	return nil
}

func (s *SoloSandbox) fnLog(args []byte) []byte {
	s.ctx.Chain.Log().Infof(string(args))
	return nil
}

func (s *SoloSandbox) fnMinted(args []byte) []byte {
	panic("implement me")
}

func (s *SoloSandbox) fnPanic(args []byte) []byte {
	s.ctx.Chain.Log().Panicf("SOLO panic: %s", string(args))
	return nil
}

func (s *SoloSandbox) fnParams(args []byte) []byte {
	return make(dict.Dict).Bytes()
}

func (s *SoloSandbox) fnPost(args []byte) []byte {
	req := wasmrequests.NewPostRequestFromBytes(args)
	if !bytes.Equal(req.ChainID.Bytes(), s.fnChainID(nil)) {
		s.Panicf("unknown chain id: %s", req.ChainID.String())
	}
	contract := s.cvt.IscpHname(req.Contract)
	if contract != iscp.Hn(s.ctx.scName) {
		s.Panicf("unknown contract: %s", contract.String())
	}
	function := s.cvt.IscpHname(req.Function)
	funcName := s.ctx.wc.FunctionFromCode(uint32(function))
	if funcName == "" {
		s.Panicf("unknown function: %s", function.String())
	}
	s.Tracef("POST %s.%s", s.ctx.scName, funcName)
	params, err := dict.FromBytes(req.Params)
	s.checkErr(err)
	scAssets := wasmlib.NewScAssetsFromBytes(req.Transfer)
	if len(scAssets) == 0 && !s.ctx.offLedger {
		s.Panicf("transfer is required for post")
	}
	if req.Delay != 0 {
		s.Panicf("cannot delay solo post")
	}
	assets := s.cvt.IscpAssets(scAssets)
	return s.postSync(s.ctx.scName, funcName, params, assets)
}

func (s *SoloSandbox) fnRequest(args []byte) []byte {
	panic("implement me")
}

func (s *SoloSandbox) fnRequestID(args []byte) []byte {
	return append(s.ctx.Chain.ChainID.Bytes()[1:], 0, 0)
}

func (s *SoloSandbox) fnResults(args []byte) []byte {
	panic("implement me")
}

// transfer tokens to L1 address
func (s *SoloSandbox) fnSend(args []byte) []byte {
	panic("implement me")
}

func (s *SoloSandbox) fnStateAnchor(args []byte) []byte {
	panic("implement me")
}

func (s *SoloSandbox) fnTimestamp(args []byte) []byte {
	return codec.EncodeInt64(time.Now().UnixNano())
}

func (s *SoloSandbox) fnTrace(args []byte) []byte {
	s.ctx.Chain.Log().Debugf(string(args))
	return nil
}

// transfer allowed tokens to L2 agent
func (s *SoloSandbox) fnTransferAllowed(args []byte) []byte {
	panic("implement me")
}
