<script lang="ts" setup>
import { mdiWeatherSunny, mdiWeatherNight } from "@mdi/js";
import { useTheme } from "vuetify";

const isDark = defineModel();
</script>

<template>
  <v-app-bar color="indigo">
    <v-app-bar-title>DNS Tools </v-app-bar-title>
    <v-switch v-model="isDark" :true-icon="mdiWeatherNight" :false-icon="mdiWeatherSunny" hide-details inset
      class="mr-4" />
  </v-app-bar>
</template>
