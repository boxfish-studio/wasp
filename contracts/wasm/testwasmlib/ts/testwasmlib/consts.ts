// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "testwasmlib";
export const ScDescription = "Exercise several aspects of WasmLib";
export const HScName       = new wasmtypes.ScHname(0x89703a45);

export const ParamAddress     = "address";
export const ParamAgentID     = "agentID";
export const ParamBlockIndex  = "blockIndex";
export const ParamBool        = "bool";
export const ParamBytes       = "bytes";
export const ParamChainID     = "chainID";
export const ParamColor       = "color";
export const ParamHash        = "hash";
export const ParamHname       = "hname";
export const ParamIndex       = "index";
export const ParamInt16       = "int16";
export const ParamInt32       = "int32";
export const ParamInt64       = "int64";
export const ParamInt8        = "int8";
export const ParamKey         = "key";
export const ParamName        = "name";
export const ParamParam       = "this";
export const ParamRecordIndex = "recordIndex";
export const ParamRequestID   = "requestID";
export const ParamString      = "string";
export const ParamUint16      = "uint16";
export const ParamUint32      = "uint32";
export const ParamUint64      = "uint64";
export const ParamUint8       = "uint8";
export const ParamValue       = "value";

export const ResultCount  = "count";
export const ResultIotas  = "iotas";
export const ResultLength = "length";
export const ResultRandom = "random";
export const ResultRecord = "record";
export const ResultValue  = "value";

export const StateArrays = "arrays";
export const StateMaps   = "maps";
export const StateRandom = "random";

export const FuncArrayAppend   = "arrayAppend";
export const FuncArrayClear    = "arrayClear";
export const FuncArraySet      = "arraySet";
export const FuncMapClear      = "mapClear";
export const FuncMapSet        = "mapSet";
export const FuncParamTypes    = "paramTypes";
export const FuncRandom        = "random";
export const FuncTakeAllowance = "takeAllowance";
export const FuncTakeBalance   = "takeBalance";
export const FuncTriggerEvent  = "triggerEvent";
export const ViewArrayLength   = "arrayLength";
export const ViewArrayValue    = "arrayValue";
export const ViewBlockRecord   = "blockRecord";
export const ViewBlockRecords  = "blockRecords";
export const ViewGetRandom     = "getRandom";
export const ViewIotaBalance   = "iotaBalance";
export const ViewMapValue      = "mapValue";

export const HFuncArrayAppend   = new wasmtypes.ScHname(0x612f835f);
export const HFuncArrayClear    = new wasmtypes.ScHname(0x88021821);
export const HFuncArraySet      = new wasmtypes.ScHname(0x2c4150b3);
export const HFuncMapClear      = new wasmtypes.ScHname(0x027f215a);
export const HFuncMapSet        = new wasmtypes.ScHname(0xf2260404);
export const HFuncParamTypes    = new wasmtypes.ScHname(0x6921c4cd);
export const HFuncRandom        = new wasmtypes.ScHname(0xe86c97ca);
export const HFuncTakeAllowance = new wasmtypes.ScHname(0x91e7bd00);
export const HFuncTakeBalance   = new wasmtypes.ScHname(0x8ad1cb27);
export const HFuncTriggerEvent  = new wasmtypes.ScHname(0xd5438ac6);
export const HViewArrayLength   = new wasmtypes.ScHname(0x3a831021);
export const HViewArrayValue    = new wasmtypes.ScHname(0x662dbd81);
export const HViewBlockRecord   = new wasmtypes.ScHname(0xad13b2f8);
export const HViewBlockRecords  = new wasmtypes.ScHname(0x16e249ea);
export const HViewGetRandom     = new wasmtypes.ScHname(0x46263045);
export const HViewIotaBalance   = new wasmtypes.ScHname(0x9d3920bd);
export const HViewMapValue      = new wasmtypes.ScHname(0x23149bef);
