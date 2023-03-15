<script lang="ts">
  import { Navbar } from '$components';
  import '../app.scss';
</script>

<Navbar />
<main class="w-full h-full flex items-center justify-center">
  <slot />
</main>
