<script>
  export let promise;
</script>

{#if promise != undefined}
  {#await promise}
    <div class="Box">
      <div class="Box-body">
        <p>Wait...</p>
      </div>
    </div>
  {:then result}
    <div class="Box">
      <div class="Box-body">
        <pre>{result}</pre>
      </div>
    </div>
  {:catch error}
    <div class="flash flash-error">
      <pre>{error}</pre>
    </div>
  {/await}
{/if}

<style lang="scss">
  pre {
    overflow-x: auto;
    overflow-y: hidden;
  }
</style>
