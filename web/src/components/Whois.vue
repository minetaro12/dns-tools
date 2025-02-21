<script lang="ts" setup>
import type { Whois } from "@/types/whois";
import { ref } from "vue";
const props = defineProps<Props>();

interface Props {
  domain: string;
}

defineExpose({
  fetchData,
});

const result = ref("");

async function fetchData() {
  if (!props.domain) {
    result.value = "ドメイン名を入力してください";
    return;
  }

  // 読み込み表示
  result.value = "Processing...";

  const body: Whois = {
    domain: props.domain,
  };

  const res = await fetch("./api/whois", {
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
  <button @click="fetchData">検索</button>
  <pre>{{ result }}</pre>
</template>

<style scoped>
button {
  padding: 10px 20px;
}

pre {
  overflow-x: auto;
  padding: 5px;
}
</style>
