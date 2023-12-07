<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  const items = [
    { label: "Whois", href: "/" },
    { label: "Query", href: "/query/" },
  ];
</script>

<div class="mt-4 ml-2">
  {#each items as item}
    <button
      class="p-2 mr-1 rounded-t hover:bg-gray-200"
      class:bg-gray-300={$page.url.pathname == item.href}
      on:click={() => goto(item.href)}
    >
      {item.label}
    </button>
  {/each}
</div>
