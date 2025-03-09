<script setup>
defineProps({
    cols: Array,
    rows: Array
})
</script>
<template>
    <table class="mt-4 min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-600">
        <thead class="bg-gray-100 dark:bg-gray-700">
            <tr>
                <th v-if="cols" v-for="col, i in cols" :key="i" v-text="col"
                    class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400" />
            </tr>
        </thead>
        <tbody v-if="rows" class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
            <tr v-for="row, i in rows" :key="i" class="hover:bg-gray-100 dark:hover:bg-gray-700">
                <td v-for="v, j in row" :key="j" v-html="v"
                    class="flex items-center p-4 mr-12 space-x-6 whitespace-nowrap dark:text-white" />
            </tr>
        </tbody>
    </table>
</template>