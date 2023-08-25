<script lang="ts">
  import { page } from "$app/stores";

  export let isOpen = false;
</script>

<div class:hidden={!isOpen} class="fixed top-14 right-4 bg-white text-black border-2 border-gray-300y px-3 py-2 rounded">
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
