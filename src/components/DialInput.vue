<script setup>
const model = defineModel({ type: String })
defineProps({
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
    id: String
})
</script>
<template>
    <div class="col-span-6">
        <label :for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{{ label }}</label>
        <input v-model="model" :name :id :type :disabled
            class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" />
    </div>
</template>