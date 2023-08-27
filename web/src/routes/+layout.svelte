<script>
  import Header from "$lib/Header.svelte";
  import Tabs from "$lib/Tabs.svelte";
  import "../app.scss";
</script>

<Header />
<div class="mx-auto max-w-[800px]">
  <Tabs />
  <div class="border-2 border-gray-300 rounded p-4">
    <slot />
  </div>
</div>
