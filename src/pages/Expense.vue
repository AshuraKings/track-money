<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from '../stores/router'
import AuthenticatedLayout from '../layouts/AuthenticatedLayout.vue'
import { authed } from '../api/withauth'
import { getExpenses } from '../api/master'
import AddExpenses from '../components/modals/AddExpenses.vue'
import DelExpense from '../components/modals/DelExpense.vue'

const expenses = ref([]), router = useRouter()

onMounted(() => {
    router.setTitle('Master Expenses')
    setTimeout(reload, 500)
})

function reload() {
    router.reverseLoading()
    authed().then(r => {
        const { body, headers, status } = r
        if (status >= 200 && status < 300) {
            router.setMenus(body.menus)
            router.setSession(body)
            router.setToken(headers.sessiontoken, headers.refreshtoken)
            if (router.path === '/master/expenses') getExpenses().then(r => {
                const { body, headers, status } = r
                if (status >= 200 && status < 300) {
                    expenses.value = body.expenses
                    router.setToken(headers.sessiontoken, headers.refreshtoken)
                } else {
                    console.log(body)
                    if (!headers.sessiontoken) {
                        router.setToken('', '')
                        router.setPath('/')
                    } else router.setToken(headers.sessiontoken, headers.refreshtoken)
                }
                router.reverseLoading()
            }).catch(e => {
                console.log(e)
                router.reverseLoading()
            })
            else router.reverseLoading()
        } else {
            console.log(body)
            if (!headers.sessiontoken) {
                router.setToken('', '')
                router.setPath('/')
            } else router.setToken(headers.sessiontoken, headers.refreshtoken)
            router.reverseLoading()
        }
    }).catch(e => {
        console.log(e)
        router.reverseLoading()
    })
}
</script>
<template>
    <AuthenticatedLayout>
        <div class="px-4 pt-6">
            <div
                class="p-4 bg-white border border-gray-200 rounded-lg shadow-sm dark:border-gray-700 sm:p-6 dark:bg-gray-800">
                <div class="items-center justify-between lg:flex">
                    <div class="mb-4 lg:mb-0">
                        <h3 class="mb-2 text-xl font-bold text-gray-900 dark:text-white">Master Expenses</h3>
                    </div>
                </div>
                <div class="sm:flex">
                    <div class="flex items-center ml-auto space-x-2 sm:space-x-3">
                        <AddExpenses @onClose="reload" />
                    </div>
                </div>
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
                                <th v-if="router.role === 'admin'"
                                    class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                                    Actions
                                </th>
                            </tr>
                        </thead>
                        <tbody class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
                            <tr v-for="e, i in expenses" :key="i" class="hover:bg-gray-100 dark:hover:bg-gray-700">
                                <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {{ i + 1 }}
                                </td>
                                <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {{ e.nm }}
                                </td>
                                <td v-if="router.role === 'admin'"
                                    class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    <DelExpense @onClose="reload" :expense="e" />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </AuthenticatedLayout>
</template>