<template>
    <div class="flex pl-0 mt-3 space-x-1 sm:pl-2 sm:mt-0">
        <div class="relative w-48 mt-1 sm:w-32 xl:w-48">
            <label class="text-gray-900 dark:text-white" :for="name">{{ label }}</label>
            <input :type :name v-model="model" :disabled :min :max
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-green-500 focus:border-green-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-green-500 dark:focus:border-green-500"
                :alt="label">
        </div>
    </div>
</template>
<script setup>
import moment from 'moment'
import { computed } from 'vue'

const model = defineModel({ type: String })
const { type } = defineProps({
    label: String,
    name: {
        required: true,
        type: String
    },
    type: {
        required: true,
        type: String
    },
    disabled: Boolean,
    min: String,
})

const max = computed(() => type === 'date' ? moment().format('YYYY-MM-DD') : undefined)
</script>