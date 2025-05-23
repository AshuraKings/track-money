<script setup>
import { computed, ref, watch } from 'vue'
import { useRouter } from '../../stores/router'
import BaseModal from '../../layouts/BaseModal.vue'
import { addMenus, getMenus } from '../../api/master'
import DialInput from '../DialInput.vue'
import DialSelect from '../DialSelect.vue'
import MessageModal from '../MessageModal.vue'

const emit = defineEmits(['onClose'])
const open = ref(false), success = ref(''), error = ref(''), label = ref(''), link = ref(''), icon = ref(''), parent = ref(''), menus = ref([])
const router = useRouter()

const invalidForm = computed(() => router.loading || !label.value)
const fixMenus = computed(() => menus.value.reduce((acc, v) => {
    acc[v.id] = v.label
    return acc
}, {}))

watch(open, (newOpen, _) => {
    if (newOpen) {
        router.reverseLoading()
        getMenus().then(r => {
            const { body, headers, status } = r
            if (status >= 200 && status < 300) {
                menus.value = body.menus
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
    }
})

function cancelled() {
    open.value = false
    success.value = error.value = label.value = link.value = icon.value = parent.value = ''
    menus.value = []
    emit('onClose')
}

function openClose() {
    open.value = !open.value
}

function submit() {
    router.reverseLoading()
    const body = { label: label.value, createdAt: "2025-03-14T21:55:54.554924Z" }
    if (link.value) body.link = link.value
    if (icon.value) body.icon = icon.value
    if (parent.value) body.parentId = parseInt(parent.value)
    addMenus(body).then(r => {
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
        class="inline-flex items-center justify-center w-1/2 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 sm:w-auto dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
        <v-icon name="md-add" />
        Add Menu
    </button>
    <BaseModal @onClose="cancelled" :open>
        <form>
            <DialInput type="text" id="label" name="label" label="Label" :disabled="router.loading" v-model="label" />
            <DialInput type="text" id="link" name="link" label="Link" :disabled="router.loading" v-model="link" />
            <DialInput type="text" id="icon" name="icon" label="Icon" :disabled="router.loading" v-model="icon" />
            <DialSelect id="parent" name="parent" label="Parent Menu" :disabled="router.loading" v-model="parent"
                :items="fixMenus" />
        </form>
        <MessageModal @onClose="cancelled" :open="success !== ''" :message="success" title="" severity="success" />
        <MessageModal @onClose="() => error = ''" :open="error !== ''" :message="error" title="" severity="danger" />
        <template #header>
            <h3 class="text-xl font-semibold dark:text-white">
                Add Menu
            </h3>
        </template>
        <template #footer>
            <button @click="cancelled" type="button" :disabled="router.loading"
                class="text-gray-900 mr-2 bg-white hover:bg-gray-100 focus:ring-4 focus:ring-primary-300 border border-gray-200 font-medium inline-flex items-center rounded-lg text-base px-3 py-2.5 text-center dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-gray-700">
                Cancel
            </button>
            <button type="submit" :disabled="invalidForm" @click.prevent="submit"
                class="text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
                Submit
            </button>
        </template>
    </BaseModal>
</template>