<script setup>
defineProps({
    options: {
        required: true,
        type: Object
    },
    value: String,
    clazz: String,
    disabled: Boolean
})
</script>
<template>
    <select @change="e => $emit('changed', e.target.value)" :disabled="disabled"
        :class="['block m-1 p-2.5 text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500', clazz]">
        <option disabled :selected="!value" value="">Pilih salah satu</option>
        <option v-for="key in Object.keys(options)" :selected="key === value" :key="key" :value="key">
            {{ options[key] }}
        </option>
    </select>
</template>