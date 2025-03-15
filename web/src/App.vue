<script setup lang="ts">
import { onMounted, ref, useTemplateRef, watch } from "vue";
import Header from "./components/Header.vue";
import { mdiGithub, mdiMagnify } from "@mdi/js";
import Whois from "./components/Whois.vue";
import Lookup from "./components/Lookup.vue";
import { useTheme } from "vuetify";

const domain = ref("");
const mode = ref<"whois" | "lookup">("whois");
const whoisRef = useTemplateRef("whoisRef");
const lookupRef = useTemplateRef("lookupRef");
const isDark = ref(window.matchMedia("(prefers-color-scheme: dark)").matches);
const theme = useTheme();

// エンターキー押下時の処理
function handleKeyDown(e: KeyboardEvent) {
  if (e.key === "Enter") {
    handleSubmit();
  }
}

// 検索ボタン押下時の処理
function handleSubmit() {
  if (mode.value === "whois") {
    whoisRef.value?.fetchData();
  } else {
    lookupRef.value?.fetchData();
  }
}

// ダークモードになっている場合はテーマを変更
onMounted(() => {
  if (isDark.value) {
    theme.global.name.value = "dark";
  }
});

// isDarkの値が変更されたらテーマを変更
watch(isDark, (value) => {
  value
    ? (theme.global.name.value = "dark")
    : (theme.global.name.value = "light");
});
</script>

<template>
  <v-app>
    <Header v-model="isDark" />
    <v-main>
      <v-container max-width="800px">
        <v-text-field
          label="Domain name"
          v-model="domain"
          placeholder="example.com"
          @keydown="handleKeyDown"
          autofocus
          clearable
        >
          <template v-slot:append>
            <v-btn
              color="indigo"
              @click="handleSubmit"
              :prepend-icon="mdiMagnify"
            >
              検索
            </v-btn>
          </template>
        </v-text-field>

        <v-tabs v-model="mode" color="indigo">
          <v-tab value="whois">Whois情報</v-tab>
          <v-tab value="lookup">DNS検索</v-tab>
        </v-tabs>
        <div v-show="mode === 'whois'">
          <Whois :domain="domain" ref="whoisRef" />
        </div>
        <div v-show="mode === 'lookup'">
          <Lookup :domain="domain" ref="lookupRef" />
        </div>
      </v-container>
    </v-main>
    <v-fab
      :icon="mdiGithub"
      href="https://github.com/minetaro12/dns-tools"
      color="indigo"
      class="position-fixed bottom-0 right-0 mb-8 mr-8"
    />
  </v-app>
</template>

<style scoped></style>
