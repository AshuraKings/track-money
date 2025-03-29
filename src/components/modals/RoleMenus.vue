<script setup>
import { ref, watch } from 'vue'
import { useRouter } from '../../stores/router'
import BaseModal from '../../layouts/BaseModal.vue'
import { addRoleMenus, getRoleMenus } from '../../api/master'
import MessageModal from '../MessageModal.vue'

const { role } = defineProps({ role: Object, })
const emit = defineEmits(['onClose'])
const open = ref(false), success = ref(''), error = ref(''), menus = ref([]), activated = ref([])
const router = useRouter()

watch(open, (newOpen, _) => {
    if (newOpen) {
        router.reverseLoading()
        getRoleMenus({ id: role.id }).then(r => {
            const { body, headers, status } = r
            if (status >= 200 && status < 300) {
                menus.value = body.menus
                activated.value = body.activatedMenus.map(m => m.id)
                router.setToken(headers.sessiontoken, headers.refreshtoken)
            } else {
                console.log(body)
                error.value = body.msg
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
    }
})

function cancelled() {
    open.value = false
    success.value = error.value = ''
    menus.value = activated.value = []
    emit('onClose')
}

function openClose() {
    open.value = !open.value
}

function submit() {
    router.reverseLoading()
    addRoleMenus({ roleId: role.id, menus: activated.value }).then(r => {
        const { body, headers, status } = r
        if (status >= 200 && status < 300) {
            success.value = body.msg
            router.setToken(headers.sessiontoken, headers.refreshtoken)
        } else {
            console.log(body)
            if (!headers.sessiontoken) {
                router.setToken('', '')
                router.setPath('/')
            } else {
                router.setToken(headers.sessiontoken, headers.refreshtoken)
                error.value = body.msg
            }
        }
        router.reverseLoading()
    }).catch(e => {
        console.log(e)
        router.reverseLoading()
    })
}
</script>
<template>
    <button @click.prevent="openClose" type="button"
        class="m-1 inline-flex items-center justify-center w-1/2 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 sm:w-auto dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
        <v-icon name="md-dashboard" />
        Role Menus
    </button>
    <BaseModal @onClose="cancelled" :open>
        <template #footer>
            <button @click="cancelled" :disabled="router.loading" type="button"
                class="text-gray-900 mr-2 bg-white hover:bg-gray-100 focus:ring-4 focus:ring-green-300 border border-gray-200 font-medium inline-flex items-center rounded-lg text-base px-3 py-2.5 text-center dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-gray-700">
                Cancel
            </button>
            <button type="submit" :disabled="router.loading" @click.prevent="submit"
                class="text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
                Submit
            </button>
        </template>
        <template #header>
            <h3 class="text-xl font-semibold dark:text-white">
                Menus Role {{ role.name }}
            </h3>
        </template>
        <MessageModal @onClose="() => { success = ''; cancelled(); }" :open="success !== ''" :message="success" title=""
            severity="success" />
        <MessageModal @onClose="() => error = ''" :open="error !== ''" :message="error" title="" severity="danger" />
        <div class="">
            <table class="mt-4 min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-600">
                <thead class="bg-gray-100 dark:bg-gray-700">
                    <tr>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Activated
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            No
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Label
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Link
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Icon
                        </th>
                        <th class="p-4 text-xs font-medium text-left text-gray-500 uppercase dark:text-gray-400">
                            Parent
                        </th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
                    <tr v-for="m, i in menus" :key="i" class="hover:bg-gray-100 dark:hover:bg-gray-700">
                        <td class="w-4 p-4">
                            <div class="flex items-center">
                                <input type="checkbox" :value="m.id" v-model="activated"
                                    class="w-4 h-4 border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:focus:ring-primary-600 dark:ring-offset-gray-800 dark:bg-gray-700 dark:border-gray-600">
                            </div>
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ i + 1 }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ m.label }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ m.link }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{ m.icon }}
                        </td>
                        <td class="p-4 text-base font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            {{m.parentId ? menus.find(m1 => m1.id === m.parentId).label : ''}}
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </BaseModal>
</template>