<script setup>
import { onMounted } from 'vue'
import AuthenticatedLayout from '../layouts/AuthenticatedLayout.vue'
import { useRouter } from '../stores/router'
import { authed, refreshToken } from '../api/withauth'

const router = useRouter()

onMounted(() => {
    router.setTitle('Dashboard')
    setTimeout(() => {
        router.reverseLoading()
        authed().then(r => {
            const { body, headers, status } = r
            if (status >= 200 && status < 300) {
                router.setMenus(body.menus)
                router.setSession(body)
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
            router.reverseLoading()
            console.log(e)
        })
    }, 500)
})
</script>
<template>
    <AuthenticatedLayout>
        <div class="py-12">
            <div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
                <div class="bg-white overflow-hidden shadow-sm sm:rounded-lg">
                    <div class="p-6 text-gray-900">You're logged in!</div>
                </div>
            </div>
        </div>
    </AuthenticatedLayout>
</template>