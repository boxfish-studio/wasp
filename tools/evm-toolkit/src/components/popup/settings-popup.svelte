<script lang="ts">
  import { Input, Select } from '$components';
  import { selectedNetwork, networks } from '../../store';

</script>

<network-settings-component class="flex flex-col space-y-6">
  {#if $selectedNetwork}
    <Select
      options={$networks}
      displayValue={network => network.text}
      bind:value={$selectedNetwork}
    />
    <Input
      id="hornetEndpoint"
      label="Hornet API endpoint"
      bind:value={$selectedNetwork.apiEndpoint}
      disabled={$selectedNetwork.id !== 1}
      stretch
    />
    <Input
      id="faucetEndpoint"
      label="Faucet API endpoint"
      bind:value={$selectedNetwork.faucetEndpoint}
      disabled={$selectedNetwork.id !== 1}
      stretch
    />
    <Input
      id="chainAddress"
      label="Chain Address"
      bind:value={$selectedNetwork.chainAddress}
      disabled={$selectedNetwork.id !== 1}
      stretch
    />
  {:else}
    <span>Loading Network Config File...</span>
  {/if}
</network-settings-component>

<style lang="scss">
</style>
