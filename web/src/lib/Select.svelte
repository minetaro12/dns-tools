<script lang="ts">
  export let value: string;
</script>

<select class="border-2 border-gray-300 bg-white rounded p-1 mt-2" bind:value>
  <slot />
</select>
