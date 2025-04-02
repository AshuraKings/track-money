<template>
    <button @click.prevent="openClose" type="button"
        class="inline-flex items-center justify-center w-1/2 px-3 py-2 text-sm font-medium text-center text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 sm:w-auto dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
        <v-icon name="md-add" />
        Add Transaction
    </button>
    <BaseModal @onClose="cancelled" :open>
        <template #header>
            <h3 class="text-xl font-semibold dark:text-white">
                Add Transaction
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
        <form>
            <DialInput type="text" id="ket" name="ket" label="Description" :disabled="router.loading" v-model="ket" />
            <DialInput type="date" id="date" name="date" label="Date" :disabled="router.loading" v-model="date" />
            <DialSelect id="fw" name="fw" label="From Wallet" :disabled="router.loading" v-model="fw"
                :items="mapWallets" />
            <DialSelect id="income" name="income" label="From Income" :disabled="router.loading" v-model="income"
                :items="mapsIncome" />
            <DialSelect id="tw" name="tw" label="To Wallet" :disabled="router.loading" v-model="tw"
                :items="mapWallets" />
            <DialSelect id="expense" name="expense" label="To Expense" :disabled="router.loading" v-model="expense"
                :items="mapsExpenses" />
            <DialInput type="number" id="amount" name="amount" label="Amount" :disabled="router.loading"
                v-model="amount" />
            <DialInput type="number" id="admin" name="admin" label="Admin Fee" :disabled="router.loading"
                v-model="admin" />
        </form>
        <MessageModal @onClose="cancelled" :open="success !== ''" :message="success" title="" severity="success" />
        <MessageModal @onClose="() => error = ''" :open="error !== ''" :message="error" title="" severity="danger" />
    </BaseModal>
</template>
<script setup>
import { computed, ref, watch } from 'vue'
import { useRouter } from '../../stores/router'
import BaseModal from '../../layouts/BaseModal.vue'
import MessageModal from '../MessageModal.vue'
import DialInput from '../DialInput.vue'
import { addTransaksies, getExpenses, getIncomes, getWallets } from '../../api/master'
import DialSelect from '../DialSelect.vue'

const emit = defineEmits(['onClose'])
const open = ref(false), success = ref(''), error = ref(''), ket = ref(''), date = ref(''), wallets = ref([]), fw = ref(0), tw = ref(0)
const incomes = ref([]), income = ref(0), expenses = ref([]), expense = ref(0), amount = ref(0), admin = ref(0)
const router = useRouter()
const mapWallets = computed(() => wallets.value.map(v => {
    const r = {}
    r[v.id] = v.nm
    return r
}).reduce((acc, m) => {
    Object.keys(m).forEach(k => acc[k] = m[k])
    return acc
}, {})), mapsIncome = computed(() => incomes.value.map(v => {
    const r = {}
    r[v.id] = v.nm
    return r
}).reduce((acc, m) => {
    Object.keys(m).forEach(k => acc[k] = m[k])
    return acc
}, {})), mapsExpenses = computed(() => expenses.value.map(v => {
    const r = {}
    r[v.id] = v.nm
    return r
}).reduce((acc, m) => {
    Object.keys(m).forEach(k => acc[k] = m[k])
    return acc
}, {})), invalidForm = computed(() => router.loading || !ket.value || !date.value || amount.value <= 0)

watch(open, (newOpen, _) => {
    if (newOpen) {
        router.reverseLoading()
        getWallets().then(r => {
            const { body, headers, status } = r
            if (status >= 200 && status < 300) {
                router.setToken(headers.sessiontoken, headers.refreshtoken)
                wallets.value = body.wallets
                getIncome1()
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
})

function getExpenses1() {
    getExpenses().then(r => {
        const { body, headers, status } = r
        if (status >= 200 && status < 300) {
            router.setToken(headers.sessiontoken, headers.refreshtoken)
            expenses.value = body.expenses
            router.reverseLoading()
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

function getIncome1() {
    getIncomes().then(r => {
        const { body, headers, status } = r
        if (status >= 200 && status < 300) {
            router.setToken(headers.sessiontoken, headers.refreshtoken)
            incomes.value = body.incomes
            getExpenses1()
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

function cancelled() {
    open.value = false
    success.value = error.value = ket.value = date.value = ''
    fw.value = tw.value = income.value = expense.value = amount.value = admin.value = 0
    wallets.value = income.value = expenses.value = []
    emit('onClose')
}

function openClose() {
    open.value = !open.value
}

function submit() {
    router.reverseLoading()
    addTransaksies({ ket: ket.value, date: date.value, amount: amount.value, admin: admin.value }).then(r => {
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