import { type Writable, writable } from 'svelte/store';
import type { NetworkOption } from '$lib/network_option';
import { SingleNodeClient, IndexerPluginClient } from '@iota/iota.js';
import { BrowserPowProvider } from '@iota/pow-browser.js';
import { persistent } from '$lib/stores';

const SELECTED_NETWORK_KEY = 'selectedNetwork';

export const networks = writable<NetworkOption[]>([]);
export const selectedNetwork: Writable<NetworkOption> = persistent(
  SELECTED_NETWORK_KEY,
  null,
);

export const indexerClient = writable<IndexerPluginClient>();
export const nodeClient = writable<SingleNodeClient>();

selectedNetwork?.subscribe(network => {
  if (!network) {
    return;
  } else {
    console.log(`Creating new client for: ${network?.apiEndpoint}`);
    const client = new SingleNodeClient(network.apiEndpoint, {
      powProvider: new BrowserPowProvider(),
    });
    nodeClient?.set(client);
    indexerClient?.set(new IndexerPluginClient(client));
  }
});
