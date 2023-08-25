<script lang="ts">
  import Menu from "./Menu.svelte";

  let isOpen = false;
  function handle() {
    isOpen = !isOpen;
  }
</script>

<header class="text-white bg-gray-800 p-4 flex justify-between items-center sticky top-0">
  <h1 class="text-2xl">
    <a href="/">DNS Tools</a>
  </h1>
  <input type="button" value="Menu" on:click={handle} class="bg-gray-800 border-2 border-gray-300 px-2 py-1 rounded hover:bg-gray-700" />
  <Menu {isOpen} />
</header>
