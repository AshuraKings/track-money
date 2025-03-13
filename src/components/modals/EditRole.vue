<script setup>
import { computed, ref, watch } from 'vue'
import { useRouter } from '../../stores/router'
import BaseModal from '../../layouts/BaseModal.vue'
import MessageModal from '../MessageModal.vue'
import DialInput from '../DialInput.vue'
import { editRole } from '../../api/master'

const { role } = defineProps({ role: Object, })
const emit = defineEmits(['onClose'])
const open = ref(false), success = ref(''), error = ref(''), nm = ref('')
const router = useRouter()

const invalidForm = computed(() => router.loading || !nm.value)

watch(open, (newOpen, _) => {
    if (newOpen) nm.value = role.name
})

function cancelled() {
    open.value = false
    success.value = error.value = nm.value = ''
    emit('onClose')
}

function openClose() {
    open.value = !open.value
}

function submit() {
    router.reverseLoading()
    editRole({ id: role.id, name: nm.value }).then(r => {
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
        class="inline-flex items-center justify-center w-1/2 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 sm:w-auto dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
        <v-icon name="md-edit" />
        Edit Role
    </button>
    <BaseModal @onClose="cancelled" :open>
        <form>
            <DialInput type="text" id="nm" name="nm" label="Name" :disabled="router.loading" v-model="nm" />
        </form>
        <MessageModal @onClose="() => { success = ''; cancelled(); }" :open="success !== ''" :message="success" title=""
            severity="success" />
        <MessageModal @onClose="() => error = ''" :open="error !== ''" :message="error" title="" severity="danger" />
        <template #header>
            <h3 class="text-xl font-semibold dark:text-white">
                Edit Role {{ role.name }}
            </h3>
        </template>
        <template #footer>
            <button @click="cancelled" :disabled="router.loading" type="button"
                class="text-gray-900 mr-2 bg-white hover:bg-gray-100 focus:ring-4 focus:ring-green-300 border border-gray-200 font-medium inline-flex items-center rounded-lg text-base px-3 py-2.5 text-center dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-gray-700">
                Cancel
            </button>
            <button type="submit" :disabled="invalidForm" @click.prevent="submit"
                class="text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
                Submit
            </button>
        </template>
    </BaseModal>
</template>