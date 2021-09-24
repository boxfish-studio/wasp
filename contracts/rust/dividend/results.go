// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableGetFactorResults struct {
	id int32
}

func (s ImmutableGetFactorResults) Factor() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxResultFactor])
}

type MutableGetFactorResults struct {
	id int32
}

func (s MutableGetFactorResults) Factor() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxResultFactor])
}

type ImmutableGetOwnerResults struct {
	id int32
}

func (s ImmutableGetOwnerResults) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxResultOwner])
}

type MutableGetOwnerResults struct {
	id int32
}

func (s MutableGetOwnerResults) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxResultOwner])
}
