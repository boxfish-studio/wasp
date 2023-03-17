import { NETWORKS } from '$lib/networks';
import type { NetworkOption } from '$lib/network_option';
import { persistent } from '$lib/stores';
import { IndexerPluginClient, SingleNodeClient } from '@iota/iota.js';
import { BrowserPowProvider } from '@iota/pow-browser.js';
import { derived, writable, type Readable, type Writable } from 'svelte/store';

const SELECTED_NETWORK_KEY = 'selectedNetworkId';
const NETWORKS_KEY = 'networks';

export const networks: Writable<NetworkOption[]> = persistent<NetworkOption[]>(NETWORKS_KEY, NETWORKS);
export const selectedNetworkId: Writable<number> = persistent(
  SELECTED_NETWORK_KEY,
  0,
);

export const selectedNetwork: Readable<NetworkOption> = derived(
  ([networks, selectedNetworkId]), ([$networks, $selectedNetworkId]) => {
    if (!$networks?.length || !($selectedNetworkId >= 0)) {
      return null;
    }
    return $networks.find(network => network.id === $selectedNetworkId);
  }
);

export function updateNetwork(network: NetworkOption) {
  networks.update($networks => {
    const index = $networks.findIndex(_network => _network?.id === network?.id);
    if (index !== -1) {
      $networks[index] = network;
    }
    return $networks;
  })
}

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
