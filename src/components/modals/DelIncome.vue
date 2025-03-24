<script setup>
import { ref } from 'vue'
import { useRouter } from '../../stores/router'
import BaseModal from '../../layouts/BaseModal.vue'
import MessageModal from '../MessageModal.vue'
import { delIncome } from '../../api/master'

const { income } = defineProps({ income: Object, })
const emit = defineEmits(['onClose'])
const open = ref(false), success = ref(''), error = ref('')
const router = useRouter()

function cancelled() {
    open.value = false
    success.value = error.value = ''
    emit('onClose')
}

function openClose() {
    open.value = !open.value
}

function submit() {
    router.reverseLoading()
    delIncome({ id: income.id }).then(r => {
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
    <BaseModal @onClose="cancelled" :open>
        <p>
            Are you sure for deleting income <b>{{ income.nm }}</b> ?
        </p>
        <template #footer>
            <button @click="cancelled" type="button" :disabled="router.loading"
                class="text-gray-900 mr-2 bg-white hover:bg-gray-100 focus:ring-4 focus:ring-red-300 border border-gray-200 font-medium inline-flex items-center rounded-lg text-base px-3 py-2.5 text-center dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-gray-700">
                Cancel
            </button>
            <button type="submit" :disabled="router.loading" @click.prevent="submit"
                class="text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-800">
                Submit
            </button>
        </template>
        <template #header>
            <h3 class="text-xl font-semibold dark:text-white">
                Delete Income {{ income.nm }}
            </h3>
        </template>
        <MessageModal @onClose="() => { success = ''; cancelled(); }" :open="success !== ''" :message="success" title=""
            severity="success" />
        <MessageModal @onClose="() => error = ''" :open="error !== ''" :message="error" title="" severity="danger" />
    </BaseModal>
    <button @click.prevent="openClose" type="button"
        class="m-1 inline-flex items-center justify-center w-1/2 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 sm:w-auto dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-800">
        <v-icon name="md-remove" />
        Delete Income
    </button>
</template>