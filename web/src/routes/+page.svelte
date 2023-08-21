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

<div class="border-2 border-gray-300 rounded mt-4 mx-auto max-w-[800px] p-4">
  <h2 class="text-xl">Whois</h2>
  <div>
    <Input bind:value={domain} placeholder="Domain" />
  </div>
  <Button {handle} />

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
