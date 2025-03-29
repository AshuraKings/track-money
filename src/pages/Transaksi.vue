<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRouter } from '../stores/router'
import AuthenticatedLayout from '../layouts/AuthenticatedLayout.vue'
import SearchInput from '../components/SearchInput.vue'
import Pagination from '../components/Pagination.vue'

const router = useRouter(), ket = ref(''), start = ref(''), end = ref(''), count = ref(100), page = ref(0)
const invalidSearch = computed(() => router.loading || (ket.value === '' && start.value === ''))

onMounted(() => {
    router.setTitle('Master Transaksi')
    setTimeout(reload, 500)
})

function reload() {
    //
}
</script>
<template>
    <AuthenticatedLayout>
        <div class="px-4 pt-6">
            <div
                class="p-4 bg-white border border-gray-200 rounded-lg shadow-sm dark:border-gray-700 sm:p-6 dark:bg-gray-800">
                <div class="items-center justify-between lg:flex">
                    <div class="mb-4 lg:mb-0">
                        <h3 class="mb-2 text-xl font-bold text-gray-900 dark:text-white">Master Transaksi</h3>
                    </div>
                </div>
                <div class="sm:flex">
                    <div class="flex items-center mb-4 sm:mb-0">
                        <SearchInput label="Description" type="text" name="ket" v-model="ket" />
                        <SearchInput label="Start Date" type="date" name="start" v-model="start" />
                        <SearchInput label="End Date" type="date" name="end" v-model="end" :disabled="!start"
                            :min="start" />
                        <div class="flex pl-0 mt-3 space-x-1 sm:pl-2 sm:mt-0">
                            <div class="relative w-48 mt-1 sm:w-64 xl:w-96">
                                <button type="button" :disabled="invalidSearch"
                                    class="inline-flex items-center justify-center w-1/2 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 sm:w-auto dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
                                    <v-icon name="fa-search" />
                                    Search
                                </button>
                            </div>
                        </div>
                    </div>
                    <div class="flex items-center ml-auto space-x-2 sm:space-x-3">
                        Opts
                    </div>
                </div>
                <div>
                    <div class="max-w-full overflow-auto">
                        <table class="mt-4 min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-600">
                            <thead class="bg-gray-100 dark:bg-gray-700">
                                <tr>
                                    <th
                                        class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                                        No
                                    </th>
                                    <th
                                        class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                                        Name
                                    </th>
                                </tr>
                            </thead>
                            <tbody class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
                                <tr class="hover:bg-gray-100 dark:hover:bg-gray-700">
                                    <td
                                        class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                        1
                                    </td>
                                    <td
                                        class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                        Sayur
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <Pagination :limit="20" :page :count @increase="() => { page += 1; reload(); }"
                        @decrease="() => { page -= 1; reload(); }" />
                </div>
            </div>
        </div>
    </AuthenticatedLayout>
</template>