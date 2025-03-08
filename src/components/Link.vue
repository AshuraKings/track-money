<script setup>
import { useRouter } from '../stores/router'

const { href } = defineProps({
    clazz: String,
    href: String,
    disabled: Boolean
})
const router = useRouter()

function go() {
    if (href) router.setPath(href)
}
</script>
<template>
    <button @click.prevent="go" type="button" :disabled="disabled" :class="clazz">
        <slot />
    </button>
</template>