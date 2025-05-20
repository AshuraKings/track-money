<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRouter } from '../stores/router'
import { authed } from '../api/withauth'
import { getChartOfMonth, getFirstDateTransaksi } from '../api/master'
import SelectSearch from '../components/SelectSearch.vue'
import moment from 'moment'

const router = useRouter(), month = ref(''), firstMonth = ref('')

const months = computed(() => {
    if (firstMonth.value !== '') {
        const bulans = {}, minMonth = moment(firstMonth.value, 'YYYY-MM-DD')
        let m = moment(), c = 3
        while (m.format('YYYY-MM') !== minMonth.format('YYYY-MM')) {
            const key = m.format('YYYY-MM'), val = m.format('MMM YYYY')
            bulans[key] = val
            m = m.subtract(1, 'months')
            if (c === 0) break
            c--
        }
        if (m.format('YYYY-MM') === minMonth.format('YYYY-MM')) {
            const key = m.format('YYYY-MM'), val = m.format('MMM YYYY')
            bulans[key] = val
        }
        return bulans
    }
    return {}
})

onMounted(() => {
    router.setTitle('Monthly Report')
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
            if (router.path === '/report/monthly') {
                if (month.value) getChartOfMonth(month.value).then(r => {
                    const { body, headers, status } = r
                    if (status >= 200 && status < 300) {
                        firstMonth.value = body.first
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
                else getFirstDateTransaksi().then(r => {
                    const { body, headers, status } = r
                    if (status >= 200 && status < 300) {
                        firstMonth.value = body.first
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
    <div class="px-4 pt-6">
        <div
            class="p-4 bg-white border border-gray-200 rounded-lg shadow-sm dark:border-gray-700 sm:p-6 dark:bg-gray-800">
            <div class="items-center justify-between lg:flex">
                <div class="mb-4 lg:mb-0">
                    <h3 class="mb-2 text-xl font-bold text-gray-900 dark:text-white">Monthly Report</h3>
                </div>
            </div>
            <div class="sm:flex">
                <div class="flex items-center mb-4 sm:mb-0">
                    <SelectSearch :disabled="router.loading" :value="month" :options="months"
                        @changed="v => { month = v; reload(); }" />
                </div>
            </div>
        </div>
    </div>
</template>