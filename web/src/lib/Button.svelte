<script lang="ts">
  export let handle: any;
</script>

<input
  type="button"
  class="bg-blue-600 mt-2 px-3 py-2 rounded text-white hover:bg-blue-500"
  value="Search"
  on:click={handle}
/>
