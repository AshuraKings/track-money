<script setup>
import { computed, onMounted, ref } from 'vue'
import NoAuthLayout from '../layouts/NoAuthLayout.vue'
import { useRouter } from '../stores/router'
import NoAuthInput from '../components/NoAuthInput.vue'
import { register } from '../api/noauth'
import MessageModal from '../components/MessageModal.vue'
import Link from '../components/Link.vue'

const router = useRouter(), nm = ref(''), username = ref(''), password = ref(''), rePass = ref(''), msg = ref(''), openMsg = ref(false)

const submitDisabled = computed(() => router.loading || !nm.value || username.value.length < 5 || password.value.length < 8 || rePass.value !== password.value)

onMounted(() => {
    router.setTitle('Register')
})

function registering() {
    router.reverseLoading()
    register({ nm: nm.value, password: password.value, username: username.value }).then(r => {
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
    <NoAuthLayout title="Register">
        <form class="mt-8 space-y-6">
            <NoAuthInput type="text" id="nm" name="nm" label="Name" :disabled="router.loading" v-model="nm" />
            <NoAuthInput type="text" id="username" name="username" label="Username" :disabled="router.loading"
                v-model="username" />
            <NoAuthInput type="password" id="password" name="password" label="Password" :disabled="router.loading"
                v-model="password" />
            <NoAuthInput type="password" id="repassword" name="repassword" label="Re-Type Password"
                :disabled="router.loading" v-model="rePass" />
            <div class="flex items-end">
                <div class="flex items-center h-5">
                    <Link href="/" :disabled="router.loading"
                        class="w-full px-4 py-2 text-base font-medium text-center text-white bg-green-700 rounded-lg hover:bg-green-800 focus:ring-4 focus:ring-green-300 sm:w-auto dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
                    To Login
                    </Link>
                </div>
            </div>
            <button @click="registering" type="button" :disabled="submitDisabled"
                class="w-full px-5 py-3 text-base font-medium text-center text-white bg-primary-700 rounded-lg hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 sm:w-auto dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
                Register
            </button>
        </form>
        <MessageModal @onClose="() => openMsg = false" :open="openMsg" title="" severity="danger" :message="msg" />
    </NoAuthLayout>
</template>