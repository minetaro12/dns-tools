<script setup lang="ts">
import { onMounted, ref, useTemplateRef } from "vue";
import Header from "./components/Header.vue";
import "./styles/global.css";
import Lookup from "./components/Lookup.vue";
import Whois from "./components/Whois.vue";

const domain = ref("");
const mode = ref<"whois" | "lookup">("whois");

function changeMode() {
  mode.value = mode.value === "whois" ? "lookup" : "whois";
}
</script>

<template>
  <Header />
  <main>
    <input type="text" placeholder="Domain name" v-model="domain" />
    <div class="mode" v-on:click="changeMode">
      <div :class="{ active: mode === 'whois' }">Whois情報</div>
      <div :class="{ active: mode === 'lookup' }">DNS検索</div>
    </div>
    <div class="child" v-show="mode === 'whois'">
      <Whois :domain="domain" ref="whoisRef" />
    </div>
    <div class="child" v-show="mode === 'lookup'">
      <Lookup :domain="domain" />
    </div>
  </main>
</template>

<style scoped>
input {
  padding: 0.5rem;
  width: 100%;
}

.mode {
  align-items: center;
  background-color: lightgray;
  border-radius: 5px;
  cursor: pointer;
  display: flex;
  gap: 1rem;
  margin: 1rem 0;
  padding: 5px;
  user-select: none;

  div {
    border-radius: 5px;
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
