<script>
  import Header from "$lib/Header.svelte";
  import Tab from "$lib/Tab.svelte";
  import "../app.scss";
</script>

<Header />
<div class="mx-auto max-w-[800px]">
  <Tab />
  <div class="border-2 border-gray-300 rounded p-4">
    <slot />
  </div>
</div>
