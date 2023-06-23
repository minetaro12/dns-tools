<script>
  import ResultBox from "./ResultBox.svelte";

  let domain;

  async function fetchData() {
    const data = new FormData();

    if (domain == "" || domain == undefined) {
      throw new Error("ドメインを入力してください");
    }

    data.append("domain", domain);

    const res = await fetch("/whois", {
      method: "POST",
      body: data,
    });
    const text = await res.text();

    if (res.ok) {
      return text;
    } else {
      throw new Error(text);
    }
  }

  let promise;
  function fetchWhois() {
    promise = fetchData();
  }
</script>

<div class="Box">
  <div class="Box-header">
    <span class="Box-title">Whois</span>
  </div>
  <div class="Box-body">
    <input
      class="form-control input-block"
      type="text"
      placeholder="example.com"
      bind:value={domain}
    />
    <button class="btn btn-outline mt-2" type="button" on:click={fetchWhois}>
      <i class="bi bi-search" />
      Search</button
    >
  </div>
</div>

<ResultBox {promise} />
