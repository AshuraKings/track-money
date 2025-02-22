<script setup>
import { computed } from 'vue'
import { useDarkMode } from './stores/darkmode'

const darkmode = useDarkMode()

const iconBtn = computed(() => darkmode.isDark ? 'md-sunny' : 'fa-moon')

function toDark() {
  darkmode.reverseTheme()
}
</script>
<template>
  <div>
    <h1 class="text-3xl font-bold underline text-black dark:text-gray-200">
      Hello world!
    </h1>
    <button @click="toDark" class="text-black dark:text-gray-200">
      <v-icon :name="iconBtn" />
      To Dark
    </button>
  </div>
</template>