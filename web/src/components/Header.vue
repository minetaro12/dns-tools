<script lang="ts" setup>
import { mdiWeatherSunny, mdiWeatherNight } from "@mdi/js";
import { ref } from "vue";
import { useTheme } from "vuetify";

const theme = useTheme();
const isDark = ref(false);

function toggleDark() {
  if (!isDark.value) {
    theme.global.name.value = "light";
  } else {
    theme.global.name.value = "dark";
  }
}
</script>

<template>
  <v-app-bar color="indigo">
    <v-app-bar-title>DNS Tools </v-app-bar-title>
    <v-switch
      v-model="isDark"
      v-on:update:model-value="toggleDark"
      :prepend-icon="isDark ? mdiWeatherNight : mdiWeatherSunny"
      hide-details
      inset
      class="mr-4"
    />
  </v-app-bar>
</template>
