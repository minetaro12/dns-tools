<script lang="ts">
  import { page } from "$app/stores";
</script>

<header class="text-white bg-gray-800 p-4 flex justify-between items-center">
  <h1 class="text-2xl">
    <a href="/">DNS Tools</a>
  </h1>
  <nav>
    <ul class="flex">
      <li>
        <a class:now={$page.url.pathname == "/"} href="/">Whois</a>
      </li>
      <li class="ml-2">
        <a class:now={$page.url.pathname == "/query/"} href="/query/">Query</a>
      </li>
    </ul>
  </nav>
</header>

<style lang="scss">
  .now {
    border-bottom: 2px solid;
  }
</style>
