<script lang="ts">
  import { page } from "$app/stores";

  let isOpen = false;
</script>

<input
  type="button"
  value="Menu"
  on:click={() => (isOpen = !isOpen)}
  class="bg-gray-800 border-2 border-gray-300 px-2 py-1 rounded cursor-pointer hover:bg-gray-700"
/>
<div
  class:hidden={!isOpen}
  class="fixed top-14 right-4 bg-white text-black border-2 border-gray-300y px-3 py-2 rounded"
>
  <nav>
    <ul class="[&>li]:m-2">
      <li>
        <a class:now={$page.url.pathname == "/"} href="/">Whois</a>
      </li>
      <li>
        <a class:now={$page.url.pathname == "/query/"} href="/query/">Query</a>
      </li>
      <li>
        <a
          href="https://api.0sn.net"
          target="_blank"
          rel="noopener noreferrer">API Tools</a
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

<style lang="scss">
  .now {
    color: royalblue;
  }
</style>
