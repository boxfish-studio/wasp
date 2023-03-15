<script lang="ts">
  import { Box, Tabs, Faucet } from '$components';
  import Withdraw from '$components/withdraw/Withdraw.svelte';

  const TABS = [
    {
      label: 'Faucet',
      value: 1,
      component: Faucet,
    },
    {
      label: 'Withdraw',
      component: Withdraw,
    },
  ];
</script>

<section class="max-w-2xl w-full">
  <Box>
    <Tabs tabs={TABS} />
  </Box>
</section>
