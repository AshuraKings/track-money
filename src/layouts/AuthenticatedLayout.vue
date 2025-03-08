<script setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { useSidebar } from '../stores/sidebar'
import AuthenticatedNavbar from '../components/AuthenticatedNavbar.vue'
import { useResponsive } from '../stores/responsive'
import AuthenticatedSidebar from '../components/AuthenticatedSidebar.vue'

const sidebar = useSidebar()
const responsive = useResponsive()
const backdrop = computed(() => 'fixed inset-0 z-10 bg-gray-900/50 dark:bg-gray-900/90 ' + (!sidebar.open ? 'hidden' : ''))

const updateWidth = () => responsive.setWidth(window.innerWidth)

onMounted(() => window.addEventListener('resize', updateWidth))
onUnmounted(() => window.removeEventListener('resize', updateWidth))
</script>
<template>
    <div>
        <AuthenticatedNavbar />
        <div class="flex pt-16 overflow-hidden bg-gray-50 dark:bg-gray-900">
            <AuthenticatedSidebar />
            <div :class="backdrop" @click.prevent="sidebar.setOpenClose" id="sidebarBackdrop"></div>
            <div class="relative w-full h-full overflow-y-auto lg:ml-64 dark:bg-gray-900 bg-smoke">
                <main>
                    <slot />
                </main>
            </div>
        </div>
    </div>
</template>