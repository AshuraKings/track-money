<script setup>
import Icon from '/vite.svg'

defineProps({
    title: {
        required: true,
        type: String
    }
})
</script>
<template>
    <div class="flex flex-col items-center justify-center px-6 pt-8 mx-auto md:h-screen pt:mt-0 dark:bg-gray-900">
        <a href="#" class="flex items-center justify-center mb-8 text-2xl font-semibold lg:mb-10 dark:text-white">
            <img :src="Icon" class="login-logo" alt="Company Logo">
        </a>
        <div class="w-full max-w-xl p-6 space-y-8 sm:p-8 bg-white rounded-lg shadow dark:bg-gray-800">
            <h2 class="text-2xl text-center font-bold text-gray-900 dark:text-white">
                {{ title }}
            </h2>
            <slot />
        </div>
    </div>
</template>