<script lang="ts">
  import Button from "$lib/Button.svelte";
  import Input from "$lib/Input.svelte";
  import Select from "$lib/Select.svelte";

  let fqdn: string, type: string, dns: string;

  async function action() {
    const res = await fetch("/lookup", {
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
  <Input bind:value={fqdn} {handle} placeholder="FQDN" />
  <Input bind:value={dns} {handle} placeholder="8.8.8.8" />
  <Select bind:value={type}>
    <option value="a">A</option>
    <option value="aaaa">AAAA</option>
    <option value="cname">CNAME</option>
    <option value="mx">MX</option>
    <option value="ns">NS</option>
    <option value="txt">TXT</option>
  </Select>
  <!-- <Input bind:value={type} {handle} placeholder="A" /> -->
</div>
<Button {handle} />

<div class="mt-4">
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
</div>
