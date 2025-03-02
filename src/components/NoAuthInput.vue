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
    <div>
        <label :for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white" v-text="label" />
        <input v-model="model" :name :id :type :disabled 
            class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" />
    </div>
</template>