<script setup>
import { useDarkMode } from './stores/darkmode'

const darkmode = useDarkMode()

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
      To Dark
    </button>
  </div>
</template>