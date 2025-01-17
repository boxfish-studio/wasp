// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

package coreroot

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type DeployContractCall struct {
	Func   *wasmlib.ScFunc
	Params MutableDeployContractParams
}

type GrantDeployPermissionCall struct {
	Func   *wasmlib.ScFunc
	Params MutableGrantDeployPermissionParams
}

type RequireDeployPermissionsCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRequireDeployPermissionsParams
}

type RevokeDeployPermissionCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRevokeDeployPermissionParams
}

type FindContractCall struct {
	Func    *wasmlib.ScView
	Params  MutableFindContractParams
	Results ImmutableFindContractResults
}

type GetContractRecordsCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetContractRecordsResults
}

type Funcs struct{}

var ScFuncs Funcs

// Deploys a non-EVM smart contract on the chain if the caller has deployment permission.
func (sc Funcs) DeployContract(ctx wasmlib.ScFuncCallContext) *DeployContractCall {
	f := &DeployContractCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncDeployContract)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// Grants deploy permission to an agent.
func (sc Funcs) GrantDeployPermission(ctx wasmlib.ScFuncCallContext) *GrantDeployPermissionCall {
	f := &GrantDeployPermissionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncGrantDeployPermission)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// Enable or disable deploy permission check
func (sc Funcs) RequireDeployPermissions(ctx wasmlib.ScFuncCallContext) *RequireDeployPermissionsCall {
	f := &RequireDeployPermissionsCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRequireDeployPermissions)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// Revokes deploy permission for an agent.
func (sc Funcs) RevokeDeployPermission(ctx wasmlib.ScFuncCallContext) *RevokeDeployPermissionCall {
	f := &RevokeDeployPermissionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRevokeDeployPermission)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// Returns the record for a given smart contract
func (sc Funcs) FindContract(ctx wasmlib.ScViewCallContext) *FindContractCall {
	f := &FindContractCall{Func: wasmlib.NewScView(ctx, HScName, HViewFindContract)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns the list of all smart contracts deployed on the chain and their records.
func (sc Funcs) GetContractRecords(ctx wasmlib.ScViewCallContext) *GetContractRecordsCall {
	f := &GetContractRecordsCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetContractRecords)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

var exportMap = wasmlib.ScExportMap{
	Names: []string{
		FuncDeployContract,
		FuncGrantDeployPermission,
		FuncRequireDeployPermissions,
		FuncRevokeDeployPermission,
		ViewFindContract,
		ViewGetContractRecords,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
	},
	Views: []wasmlib.ScViewContextFunction{
		wasmlib.ViewError,
		wasmlib.ViewError,
	},
}

func OnDispatch(index int32) {
	if index == -1 {
		exportMap.Export()
		return
	}

	panic("Calling core contract?")
}
