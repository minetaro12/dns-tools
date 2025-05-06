<script lang="ts" setup>
import type { Whois } from "@/types/whois";
import axios from "axios";
import { ref } from "vue";
const props = defineProps<Props>();
const API_URL = import.meta.env.VITE_API_URL;

interface Props {
  domain: string;
}

defineExpose({
  fetchData,
});

const result = ref("");
const isLoading = ref(false);
const errorMsg = ref("");

function fetchData() {
  isLoading.value = true;
  result.value = "";
  errorMsg.value = "";

  if (!props.domain) {
    errorMsg.value = "ドメイン名を入力してください";
    isLoading.value = false;
    return;
  }

  const body: Whois = {
    domain: props.domain,
  };

  axios
    .post(`${API_URL}/api/whois`, body)
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
  <v-alert type="error" class="mt-2" :text="errorMsg" v-if="errorMsg != ''" />
  <v-progress-linear indeterminate class="mt-2" v-if="isLoading" />
  <v-card variant="tonal" class="mt-2">
    <pre class="pa-3 overflow-x-auto text-body-2" v-if="result">{{
      result
    }}</pre>
  </v-card>
</template>

<style scoped></style>
