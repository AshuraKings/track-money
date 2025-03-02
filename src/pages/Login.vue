<script setup>
import { onMounted, ref } from 'vue'
import NoAuthInput from '../components/NoAuthInput.vue'
import NoAuthLayout from '../layouts/NoAuthLayout.vue'
import { useRouter } from '../stores/router'
import { login } from '../api/noauth'
import MessageModal from '../components/MessageModal.vue'

const username = ref(''), password = ref(''), openMsg = ref(false), msg = ref('')
const router = useRouter()

onMounted(() => {
    router.setTitle('Login')
})

function loggingIn() {
    router.reverseLoading()
    login({ username: username.value, password: password.value }).then(r => {
        const { body, headers, status } = r
        if (status >= 200 && status < 300) {
            router.setToken(headers.sessiontoken, headers.refreshtoken)
            router.setPath('/dashboard')
        } else {
            msg.value = body.msg
            openMsg.value = true
        }
        router.reverseLoading()
    }).catch(e => {
        router.reverseLoading()
        console.log(e)
    })
}
</script>
<template>
    <NoAuthLayout title="Login">
        <form class="mt-8 space-y-6">
            <NoAuthInput type="text" id="username" name="username" label="Username" :disabled="router.loading"
                v-model="username" />
            <NoAuthInput type="password" id="password" name="password" label="Password" :disabled="router.loading"
                v-model="password" />
            <div class="flex items-end">
                <div class="flex items-center h-5">
                    <button @click="() => router.setPath('/register')" type="button" :disabled="router.loading"
                        class="w-full px-4 py-2 text-base font-medium text-center text-white bg-green-700 rounded-lg hover:bg-green-800 focus:ring-4 focus:ring-green-300 sm:w-auto dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
                        To Register
                    </button>
                </div>
            </div>
            <button @click="loggingIn" type="button" :disabled="router.loading"
                class="w-full px-5 py-3 text-base font-medium text-center text-white bg-primary-700 rounded-lg hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 sm:w-auto dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
                Login
            </button>
        </form>
        <MessageModal @onClose="() => openMsg = false" :open="openMsg" title="" severity="danger" :message="msg" />
    </NoAuthLayout>
</template>