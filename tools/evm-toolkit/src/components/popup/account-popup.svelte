<script lang="ts">
  import {
    truncateText,
    handleEnterKeyDown,
    copyToClipboard,
  } from '$lib/utils';
  import { Button } from '..';
  import Tooltip from '../tooltip.svelte';

  export let account = undefined;
  let showCopiedTooltip = false;
  const handleCopyToClipboard = (copyValue: string) => {
    copyToClipboard(copyValue);
    showCopiedTooltip = true;
  };
</script>

<div class="flex flex-row items-center py-6">
  <account-box class:connected={account}>
    {#if account}
    <account-address class="flex items-center justify-between w-full">
      <div class="dot-primary flex items-center justify-between space-x-2">
        <span class="font-semibold">{truncateText(account)}</span>
        <Tooltip message="Copied" show={showCopiedTooltip}>
          <img
            src="/copy-icon.svg"
            alt="Copy address"
            on:click={() => handleCopyToClipboard(account)}
            on:keydown={event =>
              handleEnterKeyDown(event, () => handleCopyToClipboard(account))}
          />
        </Tooltip>
      </div>
      <Button
        title="Disconnect"
        outline
        compact
        on:click={() => {
          console.log('disconnect');
        }}
      />
    </account-address>
    {:else}
      <metamask-box class="flex items-center justify-center space-x-4">
        <img src="/metamask-logo.png" alt="metamask icon" />
        <span class="text-white font-semibold">Connect your wallet</span>
      </metamask-box>
    {/if}
  </account-box>
</div>

<style lang="scss">
  account-box {
    @apply relative;
    @apply w-full;
    @apply flex;
    @apply justify-between;
    @apply bg-shimmer-background-tertiary;
    @apply rounded-xl;
    @apply p-4;
    @apply cursor-pointer;
    @apply hover:bg-shimmer-background-tertiary-hover;
    @apply transition-all;
    @apply duration-200;
    @apply ease-in-out;
    &.connected {
      @apply pl-8;
    }
  }
</style>
