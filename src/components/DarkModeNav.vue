<script setup>
import { useDarkMode } from '../stores/darkmode'

const darkmode = useDarkMode()
</script>
<template>
    <button @click.prevent="darkmode.reverseTheme" type="button"
        class="text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 rounded-lg text-sm p-2.5">
        <v-icon class="w-5 h-5" :name="darkmode.isDark ? 'md-sunny' : 'fa-moon'" />
    </button>
</template>