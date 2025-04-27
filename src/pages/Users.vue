<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from '../stores/router'
import { getUsers } from '../api/master'
import AddUser from '../components/modals/AddUser.vue'
import EditUser from '../components/modals/EditUser.vue'
import DelUser from '../components/modals/DelUser.vue'
import { authed } from '../api/withauth'

const router = useRouter()
const users = ref([])

onMounted(() => {
    router.setTitle('Master Users')
    setTimeout(() => reload(), 500)
})

function reload() {
    router.reverseLoading()
    authed().then(r => {
        const { body, headers, status } = r
        if (status >= 200 && status < 300) {
            router.setMenus(body.menus)
            router.setSession(body)
            router.setToken(headers.sessiontoken, headers.refreshtoken)
            if (router.path === '/master/users') getUsers().then(r => {
                const { body, headers, status } = r
                if (status >= 200 && status < 300) {
                    users.value = body.users
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
        router.reverseLoading()
        console.log(e)
    })
}
</script>
<template>
    <div class="px-4 pt-6">
        <div
            class="p-4 bg-white border border-gray-200 rounded-lg shadow-sm dark:border-gray-700 sm:p-6 dark:bg-gray-800">
            <div class="items-center justify-between lg:flex">
                <div class="mb-4 lg:mb-0">
                    <h3 class="mb-2 text-xl font-bold text-gray-900 dark:text-white">Master Users</h3>
                </div>
            </div>
            <div class="sm:flex">
                <div class="flex items-center ml-auto space-x-2 sm:space-x-3">
                    <AddUser @onClose="reload" />
                </div>
            </div>
            <table class="mt-4 min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-600">
                <thead class="bg-gray-100 dark:bg-gray-700">
                    <tr>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            No
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Name
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Username
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Role
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Actions
                        </th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
                    <tr v-for="user, i in users" :key="i" class="hover:bg-gray-100 dark:hover:bg-gray-700">
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ i + 1 }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ user.name }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ user.username }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ user.role }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            <EditUser @onClose="reload" :user />
                            <DelUser @onClose="reload" :user />
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>