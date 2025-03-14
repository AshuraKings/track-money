<script setup>
import BaseModal from '../layouts/BaseModal.vue'

defineProps({
    open: Boolean,
    title: String,
    message: String,
    severity: {
        validator(value, props) {
            return ['success', 'warning', 'danger'].includes(value)
        }
    }
})
</script>
<template>
    <BaseModal @onClose="$emit('onClose')" :open="open">
        <template #header>
            <h3 class="text-xl font-semibold dark:text-white">{{ title }}</h3>
        </template>
        <div class="text-center">
            <v-icon v-if="severity === 'success'" fill="blue" name="io-checkmark-sharp"
                class="w-32 h-32 text-primary-500 fill-primary-500" />
            <v-icon v-else name="md-cancel-outlined" :fill="severity === 'danger' ? 'red' : 'yellow'"
                :class="`w-32 h-32 text-${severity === 'danger' ? 'red' : 'yellow'}-500 fill-${severity === 'danger' ? 'red' : 'yellow'}-500`" />
            <p class="text-black dark:text-white">{{ message }}</p>
        </div>
    </BaseModal>
</template>