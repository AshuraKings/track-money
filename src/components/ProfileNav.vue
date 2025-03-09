<script setup>
import { computed, ref } from 'vue'
import { useResponsive } from '../stores/responsive'
import { useRouter } from '../stores/router'
import { logout } from '../api/withauth'

const isClose = ref(true), responsive = useResponsive(), router = useRouter()
const styleDropdown = computed(() => 'position: absolute; inset: 0px auto auto 0px; margin: 0px; transform: translate(' + (responsive.windowWidth - 150) + 'px, 61px);')

function out() {
    router.reverseLoading()
    logout().then(r => {
        const { body, status } = r
        if (status >= 200 && status < 300) {
            router.setToken('', '')
            router.setPath('/')
        } else {
            console.log(body)
            if (body.msg === 'Token is expired') {
                router.setToken('', '')
                router.setPath('/')
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
    <div class="flex items-center ml-3">
        <div>
            <button @click.prevent="() => isClose = !isClose" type="button"
                class="flex text-sm bg-gray-800 rounded-full focus:ring-4 focus:ring-gray-300 dark:focus:ring-gray-600">
                <span class="sr-only">Open user menu</span>
                <img class="w-8 h-8 rounded-full" src="https://flowbite.com/docs/images/people/profile-picture-5.jpg"
                    alt="user photo">
            </button>
        </div>
        <div v-if="!isClose" :style="styleDropdown"
            class="z-50 my-4 text-base list-none bg-white divide-y divide-gray-100 rounded shadow dark:bg-gray-700 dark:divide-gray-600 block">
            <div class="px-4 py-3">
                <p class="text-sm text-gray-900 dark:text-white" v-text="router.name" />
                <p class="text-sm font-medium text-gray-900 truncate dark:text-gray-300" v-text="router.username" />
            </div>
            <ul class="py-1">
                <li>
                    <button type="button" @click.prevent="() => router.setPath('/dashboard')"
                        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-600 dark:hover:text-white">
                        Dashboard
                    </button>
                </li>
                <li>
                    <button type="button" @click.prevent="() => router.setPath('/profile')"
                        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-600 dark:hover:text-white">
                        Settings
                    </button>
                </li>
                <li>
                    <button type="button" @click.prevent="out"
                        class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-600 dark:hover:text-white">
                        Logout
                    </button>
                </li>
            </ul>
        </div>
    </div>
</template>