<script lang="ts" setup>
import { mdiWeatherSunny, mdiWeatherNight } from "@mdi/js";
import { useTheme } from "vuetify";

const theme = useTheme();
const isDark = defineModel();
</script>

<template>
  <v-app-bar color="indigo">
    <v-app-bar-title>DNS Tools </v-app-bar-title>
    <v-switch
      v-model="isDark"
      :prepend-icon="isDark ? mdiWeatherNight : mdiWeatherSunny"
      hide-details
      inset
      class="mr-4"
    />
  </v-app-bar>
</template>
