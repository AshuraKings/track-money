<template>
    <div
        class="sticky bottom-0 right-0 items-center w-full p-4 bg-white border-t border-gray-200 sm:flex sm:justify-between dark:bg-gray-800 dark:border-gray-700">
        <div class="flex items-center mb-4 sm:mb-0">
            <button type="button" @click="$emit('decrease')" :disabled="!page"
                class="inline-flex justify-center p-1 text-gray-500 rounded cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                <v-icon name="fa-chevron-left" />
            </button>
            <button type="button" @click="$emit('increase')" :disabled="((page + 1) * limit) >= count"
                class="inline-flex justify-center p-1 text-gray-500 rounded cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                <v-icon name="fa-chevron-right" />
            </button>
            <span class="text-sm font-normal text-gray-500 dark:text-gray-400">
                Showing
                <span class="font-semibold text-gray-900 dark:text-white">
                    {{ (page * limit) + 1 }}-{{ (page + 1) * limit }}
                </span>
                of <span class="font-semibold text-gray-900 dark:text-white">{{ count }}</span>
            </span>
        </div>
        <div class="flex items-center space-x-3">
            <button type="button" @click="$emit('decrease')" :disabled="!page"
                class="inline-flex items-center justify-center flex-1 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
                <v-icon name="fa-chevron-left" />
                Previous
            </button>
            <button type="button" @click="$emit('increase')" :disabled="((page + 1) * limit) >= count"
                class="inline-flex items-center justify-center flex-1 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
                Next
                <v-icon name="fa-chevron-right" />
            </button>
        </div>
    </div>
</template>
<script setup>
defineProps({
    count: {
        type: Number,
        required: true
    },
    limit: {
        type: Number,
        required: true
    },
    page: {
        type: Number,
        required: true
    },
})
</script>