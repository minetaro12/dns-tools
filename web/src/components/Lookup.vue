<script lang="ts" setup>
import type { Lookup } from "@/types/lookup";
import axios from "axios";
import { ref } from "vue";
const props = defineProps<Props>();

interface Props {
  domain: string;
}

defineExpose({
  fetchData,
});

const result = ref("");
const isLoading = ref(false);
const errorMsg = ref("");
const type = ref("A");
const dns = ref("8.8.8.8");

function fetchData() {
  isLoading.value = true;
  result.value = "";
  errorMsg.value = "";

  if (!props.domain) {
    errorMsg.value = "ドメイン名を入力してください";
    isLoading.value = false;
    return;
  }

  const body: Lookup = {
    fqdn: props.domain,
    dns: dns.value,
    type: type.value,
  };

  axios
    .post("./api/lookup", body)
    .then((res) => {
      result.value = res.data;
    })
    .catch((err) => {
      errorMsg.value = err.response.data;
    })
    .finally(() => {
      isLoading.value = false;
    });
}
</script>

<template>
  <v-row class="mt-0">
    <v-col cols="4">
      <v-select
        v-model="type"
        :items="['A', 'AAAA', 'CNAME', 'MX', 'NS', 'TXT']"
        label="Type"
      />
    </v-col>
    <v-col cols="8">
      <v-text-field label="DNS" v-model="dns" placeholder="8.8.8.8" clearable />
    </v-col>
  </v-row>
  <v-alert type="error" :text="errorMsg" v-if="errorMsg != ''" />
  <v-progress-linear indeterminate v-if="isLoading" />
  <v-card variant="tonal">
    <pre class="pa-3 overflow-x-auto text-body-2" v-if="result">{{
      result
    }}</pre>
  </v-card>
</template>

<style scoped>
pre {
  overflow-x: auto;
  padding: 5px;
}

.input-group {
  display: flex;
  gap: 10px;

  @media screen and (max-width: 600px) {
    flex-direction: column;
  }
}

input,
select {
  padding: 10px;
}
</style>
