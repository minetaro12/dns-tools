<script lang="ts">
  import Button from "$lib/Button.svelte";
  import Input from "$lib/Input.svelte";

  let domain: string;

  async function action() {
    const res = await fetch("/whois", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ domain: domain }),
    });
    const result = await res.text();

    if (res.ok) {
      return result;
    } else {
      throw new Error(result);
    }
  }

  let promise: Promise<String> | null = null;

  function handle() {
    promise = action();
  }
</script>

<svelte:head>
  <title>DNS Tools - Whois</title>
</svelte:head>

<div>
  <Input bind:value={domain} {handle} placeholder="Domain" />
</div>
<Button {handle} />

<div class="mt-4">
  {#await promise}
    <p>Wait...</p>
  {:then result}
    {#if result != null}
      <pre class="text-sm overflow-x-auto">{result}</pre>
    {/if}
  {:catch error}
  <pre class="text-sm text-red-600 overflow-x-auto">Error</pre>
    <pre class="text-sm text-red-600 overflow-x-auto">{error.message}</pre>
  {/await}
</div>
