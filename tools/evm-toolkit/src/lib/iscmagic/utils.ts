import { Blake2b } from '@iota/crypto.js';
import { Converter } from '@iota/util.js';
import { Buffer } from 'buffer';

import { SimpleBufferCursor } from '$lib/simple-buffer-cursor';

export function evmAddressToAgentID(evmStoreAccount: string): Uint8Array {
  // This function constructs an AgentID that is required to be used with contracts
  // Wasp understands different AgentID types and each AgentID needs to provide a certain ID that describes it's address type.
  // In the case of EVM addresses it's ID 3.
  const agentIDKindEthereumAddress = 3;

  const receiverAddrBinary = Converter.hexToBytes(evmStoreAccount);
  const addressBytes = new Uint8Array([
    agentIDKindEthereumAddress,
    ...receiverAddrBinary,
  ]);

  return addressBytes;
}

export function hNameFromString(name): number {
  const ScHNameLength = 4;
  const stringBytes = Converter.utf8ToBytes(name);
  const hash = Blake2b.sum256(stringBytes);

  for (let i = 0; i < hash.length; i += ScHNameLength) {
    const slice = hash.slice(i, i + ScHNameLength);
    const cursor = new SimpleBufferCursor(Buffer.from(slice));

    return cursor.readUInt32LE();
  }

  return 0;
}
