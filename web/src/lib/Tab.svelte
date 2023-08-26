<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
</script>

<div class="mt-4 ml-2">
  <button
    class="p-2 rounded-t"
    class:bg-gray-300={$page.url.pathname === "/"}
    on:click={() => goto("/")}
  >
    Whois
  </button>
  <button
    class="p-2 rounded-t"
    class:bg-gray-300={$page.url.pathname === "/query/"}
    on:click={() => goto("/query/")}>Query</button
  >
</div>
