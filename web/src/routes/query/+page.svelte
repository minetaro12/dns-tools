<script lang="ts">
  import Button from "$lib/Button.svelte";
  import Input from "$lib/Input.svelte";

  let fqdn: string, type: string, dns: string;

  async function action() {
    const res = await fetch("/nslookup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ fqdn: fqdn, type: type, dns: dns }),
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
  <title>DNS Tools - Query</title>
</svelte:head>

<div>
  <Input bind:value={fqdn} placeholder="FQDN" />
  <Input bind:value={dns} placeholder="8.8.8.8" />
  <Input bind:value={type} placeholder="A" />
</div>
<Button {handle} />

<div class="mt-4">
  <div class="mt-4">
    {#await promise}
      <p>Wait...</p>
    {:then result}
      {#if result != null}
        <pre class="text-xs overflow-auto">{result}</pre>
      {/if}
    {:catch error}
      <p class="text-red-700">{error.message}</p>
    {/await}
  </div>
</div>
