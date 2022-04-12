// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as wasmtypes from "wasmlib/wasmtypes";

export class Erc721Events {

	approval(
		approved: wasmtypes.ScAgentID,
		owner: wasmtypes.ScAgentID,
		tokenID: wasmtypes.ScHash,
	): void {
		const evt = new wasmlib.EventEncoder("erc721.approval");
		evt.encode(wasmtypes.agentIDToString(approved));
		evt.encode(wasmtypes.agentIDToString(owner));
		evt.encode(wasmtypes.hashToString(tokenID));
		evt.emit();
	}

	approvalForAll(
		approval: bool,
		operator: wasmtypes.ScAgentID,
		owner: wasmtypes.ScAgentID,
	): void {
		const evt = new wasmlib.EventEncoder("erc721.approvalForAll");
		evt.encode(wasmtypes.boolToString(approval));
		evt.encode(wasmtypes.agentIDToString(operator));
		evt.encode(wasmtypes.agentIDToString(owner));
		evt.emit();
	}

	init(
		name: string,
		symbol: string,
	): void {
		const evt = new wasmlib.EventEncoder("erc721.init");
		evt.encode(wasmtypes.stringToString(name));
		evt.encode(wasmtypes.stringToString(symbol));
		evt.emit();
	}

	mint(
		balance: u64,
		owner: wasmtypes.ScAgentID,
		tokenID: wasmtypes.ScHash,
	): void {
		const evt = new wasmlib.EventEncoder("erc721.mint");
		evt.encode(wasmtypes.uint64ToString(balance));
		evt.encode(wasmtypes.agentIDToString(owner));
		evt.encode(wasmtypes.hashToString(tokenID));
		evt.emit();
	}

	transfer(
		from: wasmtypes.ScAgentID,
		to: wasmtypes.ScAgentID,
		tokenID: wasmtypes.ScHash,
	): void {
		const evt = new wasmlib.EventEncoder("erc721.transfer");
		evt.encode(wasmtypes.agentIDToString(from));
		evt.encode(wasmtypes.agentIDToString(to));
		evt.encode(wasmtypes.hashToString(tokenID));
		evt.emit();
	}
}
