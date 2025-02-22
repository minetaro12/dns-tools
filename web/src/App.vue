<script setup lang="ts">
import { ref, useTemplateRef } from "vue";
import Header from "./components/Header.vue";
import "./styles/global.css";
import Lookup from "./components/Lookup.vue";
import Whois from "./components/Whois.vue";

const domain = ref("");
const mode = ref<"whois" | "lookup">("whois");
const whoisRef = useTemplateRef("whoisRef");
const lookupRef = useTemplateRef("lookupRef");

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
</script>

<template>
  <Header />
  <main>
    <div class="input-group">
      <input
        type="text"
        placeholder="Domain name"
        v-model="domain"
        v-on:keydown="handleKeyDown"
        autofocus
      />
      <button v-on:click="handleSubmit">検索</button>
    </div>
    <div class="mode">
      <div :class="{ active: mode === 'whois' }" v-on:click="mode = 'whois'">
        Whois情報
      </div>
      <div :class="{ active: mode === 'lookup' }" v-on:click="mode = 'lookup'">
        DNS検索
      </div>
    </div>
    <div class="child" v-show="mode === 'whois'">
      <Whois :domain="domain" ref="whoisRef" />
    </div>
    <div class="child" v-show="mode === 'lookup'">
      <Lookup :domain="domain" ref="lookupRef" />
    </div>
  </main>
</template>

<style scoped>
.input-group {
  display: flex;
  gap: 10px;

  button {
    width: 80px;
  }
}

input {
  padding: 0.5rem;
  width: 100%;
}

.mode {
  align-items: center;
  background-color: lightgray;
  border-radius: 5px;
  display: flex;
  gap: 1rem;
  margin: 1rem 0;
  padding: 5px;
  user-select: none;

  div {
    border-radius: 5px;
    cursor: pointer;
    flex: 1;
    text-align: center;
    line-height: 1.4;
    padding: 5px;
  }

  div.active {
    background-color: white;
  }
}
</style>
