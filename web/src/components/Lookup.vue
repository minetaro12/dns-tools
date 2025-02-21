<script lang="ts" setup>
import type { Lookup } from "@/types/lookup";
import { ref } from "vue";
const props = defineProps<Props>();

interface Props {
  domain: string;
}

defineExpose({
  fetchData,
});

const result = ref("");
const type = ref("a");
const dns = ref("");

async function fetchData() {
  if (!props.domain) {
    result.value = "ドメイン名を入力してください";
    return;
  }

  // 読み込み表示
  result.value = "Processing...";

  const body: Lookup = {
    fqdn: props.domain,
    dns: dns.value,
    type: type.value,
  };

  const res = await fetch("./api/lookup", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  if (!res.ok) {
    result.value = "Error fetching data";
  } else [(result.value = await res.text())];
}
</script>

<template>
  <div class="input-group">
    <select v-model="type">
      <option value="a">A</option>
      <option value="aaaa">AAAA</option>
      <option value="cname">CNAME</option>
      <option value="mx">MX</option>
      <option value="ns">NS</option>
      <option value="txt">TXT</option>
    </select>
    <input type="text" v-model="dns" placeholder="8.8.8.8" />
    <button @click="fetchData">検索</button>
  </div>
  <pre>{{ result }}</pre>
</template>

<style scoped>
pre {
  overflow-x: auto;
  padding: 5px;
}

.input-group {
  display: flex;
  gap: 10px;
}

input,
select {
  padding: 10px;
}
button {
  padding: 10px 20px;
}
</style>
