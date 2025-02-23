<script setup>
import { computed, onMounted, ref } from 'vue'
import { useDarkMode } from './stores/darkmode'
import NoAuthLayout from './layouts/NoAuthLayout.vue'

const darkmode = useDarkMode(), msg = ref('')

const iconBtn = computed(() => darkmode.isDark ? 'md-sunny' : 'fa-moon')

function toDark() {
  darkmode.reverseTheme()
}

onMounted(() => {
  document.title = 'App'
  fetch('/api', { method: 'GET' })
    .then(r => r.json())
    .then(r => msg.value = r.msg)
    .catch(console.log)
})
</script>
<template>
  <NoAuthLayout title="No Auth Test">
    <h1 class="text-3xl font-bold underline text-black dark:text-gray-200">
      {{ msg }}
    </h1>
    <button @click="toDark" class="text-black dark:text-gray-200">
      <v-icon :name="iconBtn" />
      To Dark
    </button>
  </NoAuthLayout>
</template>