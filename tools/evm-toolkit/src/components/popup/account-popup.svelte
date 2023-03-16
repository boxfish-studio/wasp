<script lang="ts">
  import {
    truncateText,
    handleEnterKeyDown,
    copyToClipboard,
  } from '$lib/utils';

  export let account;
  const handleCopyToClipboard = (copyValue: string) => {
    copyToClipboard(copyValue);
  };
</script>

<div class="flex flex-row items-center py-6">
  <account-box
    class="w-full flex justify-between bg-shimmer-background-tertiary rounded-xl p-4 pl-8 relative"
  >
    <account-address
      class="dot-primary flex items-center justify-between space-x-2"
    >
      <span class="font-semibold">{truncateText(account)}</span>
      <img
        src="/copy-icon.svg"
        alt="Copy address"
        on:click={() => handleCopyToClipboard(account)}
        on:keydown={event =>
          handleEnterKeyDown(event, () => handleCopyToClipboard(account))}
      />
    </account-address>
  </account-box>
</div>
