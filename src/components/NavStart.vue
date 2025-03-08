<script setup>
import { computed } from 'vue'
import { useSidebar } from '../stores/sidebar'
import Icon from '/vite.svg'

const sidebar = useSidebar()

const toggleSidebarMobileClose = computed(() => 'w-6 h-6 ' + (!sidebar.open ? 'hidden' : ''))
const toggleSidebarMobile = computed(() => 'w-6 h-6 ' + (sidebar.open ? 'hidden' : ''))
</script>
<template>
    <div class="flex items-center justify-start">
        <button @click.prevent="sidebar.setOpenClose" type="button"
            class="p-2 text-gray-600 rounded cursor-pointer lg:hidden hover:text-gray-900 hover:bg-gray-100 focus:bg-gray-100 dark:focus:bg-gray-700 focus:ring-2 focus:ring-gray-100 dark:focus:ring-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
            <svg :class="toggleSidebarMobile" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                    d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                    clip-rule="evenodd"></path>
            </svg>
            <svg :class="toggleSidebarMobileClose" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                    d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                    clip-rule="evenodd"></path>
            </svg>
        </button>
        <a href="#" class="flex ml-2 md:mr-24">
            <img :src="Icon" class="h-8 mr-3" alt="Company Logo">
            <span class="self-center text-xl font-semibold sm:text-2xl whitespace-nowrap dark:text-white">Company</span>
        </a>
    </div>
</template>