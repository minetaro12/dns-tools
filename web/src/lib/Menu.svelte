<script lang="ts">
  let isOpen = false;
</script>

<button
  on:click={() => (isOpen = !isOpen)}
  class="bg-gray-800 border-2 border-gray-300 px-2 py-1 rounded hover:bg-gray-700"
  >Menu</button
>
<div
  class:hidden={!isOpen}
  class="fixed top-14 right-4 bg-white text-black border-2 border-gray-300y px-3 py-2 rounded"
>
  <nav>
    <ul class="[&>li]:m-2">
      <li>
        <a href="https://api.0sn.net" target="_blank" rel="noopener noreferrer"
          >API Tools</a
        >
      </li>
      <li>
        <a
          href="https://github.com/minetaro12/dns-tools"
          target="_blank"
          rel="noopener noreferrer">GitHub</a
        >
      </li>
    </ul>
  </nav>
</div>
