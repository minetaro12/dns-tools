<script>
  import ResultBox from "./ResultBox.svelte";
  let fqdn, qType, dns;

  async function fetchData(digType) {
    const data = new FormData();

    if (fqdn == "" || fqdn == undefined) {
      throw new Error("FQDNを入力してください");
    }

    if (qType == "" || qType == undefined) {
      qType = "";
    }

    if (dns == "" || dns == undefined) {
      dns = "";
    }

    data.append("fqdn", fqdn);
    data.append("type", qType);
    data.append("dns", dns);

    const res = await fetch(`/${digType}`, {
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
  function fetchDig() {
    promise = fetchData("dig");
  }

  function fetchNslookup() {
    promise = fetchData("nslookup");
  }
</script>

<div class="Box">
  <div class="Box-header">
    <span class="Box-title">Query</span>
  </div>
  <div class="Box-body">
    <input
      class="form-control input-block"
      bind:value={dns}
      placeholder="8.8.8.8"
    />
    <input
      class="form-control input-block mt-2"
      bind:value={fqdn}
      placeholder="example.com"
    />
    <input
      class="form-control input-block mt-2"
      bind:value={qType}
      placeholder="A"
    />
    <button class="btn btn-outline mt-2" type="button" on:click={fetchDig}>
      <i class="bi bi-search" />
      dig
    </button>
    <button
      class="btn btn-outline mt-2 ml-2"
      type="button"
      on:click={fetchNslookup}
    >
      <i class="bi bi-search" />
      nslookup
    </button>
  </div>
</div>

<ResultBox {promise} />
