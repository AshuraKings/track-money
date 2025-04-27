<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from '../stores/router'
import { authed } from '../api/withauth'
import { getRoles } from '../api/master'
import AddRole from '../components/modals/AddRole.vue'
import EditRole from '../components/modals/EditRole.vue'
import DelRole from '../components/modals/DelRole.vue'
import RoleMenus from '../components/modals/RoleMenus.vue'

const router = useRouter(), roles = ref([])

onMounted(() => {
    router.setTitle('Master Roles')
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
            if (router.path === '/master/roles') getRoles().then(r => {
                const { body, headers, status } = r
                if (status >= 200 && status < 300) {
                    roles.value = body.roles
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
                    <h3 class="mb-2 text-xl font-bold text-gray-900 dark:text-white">Master Roles</h3>
                </div>
            </div>
            <div class="sm:flex">
                <div class="flex items-center ml-auto space-x-2 sm:space-x-3">
                    <AddRole @onClose="reload" />
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
                            Actions
                        </th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
                    <tr v-for="r, i in roles" :key="i" class="hover:bg-gray-100 dark:hover:bg-gray-700">
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ i + 1 }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ r.name }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            <EditRole @onClose="reload" :role="r" />
                            <RoleMenus @onClose="reload" :role="r" />
                            <DelRole @onClose="reload" :role="r" />
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>