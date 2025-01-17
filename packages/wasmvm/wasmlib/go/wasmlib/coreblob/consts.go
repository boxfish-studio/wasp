// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

package coreblob

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "blob"
	ScDescription = "Blob Contract"
	HScName       = wasmtypes.ScHname(0xfd91bc63)
)

const (
	ParamBlobs      = "this"
	ParamDataSchema = "d"
	ParamField      = "field"
	ParamHash       = "hash"
	ParamProgBinary = "p"
	ParamSources    = "s"
	ParamVMType     = "v"
)

const (
	ResultBlobSizes = "this"
	ResultBytes     = "bytes"
	ResultHash      = "hash"
)

const (
	FuncStoreBlob    = "storeBlob"
	ViewGetBlobField = "getBlobField"
	ViewGetBlobInfo  = "getBlobInfo"
	ViewListBlobs    = "listBlobs"
)

const (
	HFuncStoreBlob    = wasmtypes.ScHname(0xddd4c281)
	HViewGetBlobField = wasmtypes.ScHname(0x1f448130)
	HViewGetBlobInfo  = wasmtypes.ScHname(0xfde4ab46)
	HViewListBlobs    = wasmtypes.ScHname(0x62ca7990)
)
