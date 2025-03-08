<script setup>
import { computed, ref } from 'vue'
import { useSidebar } from '../stores/sidebar'
import { useRouter } from '../stores/router'
import SidebarMenu from './SidebarMenu.vue'

const sidebar = useSidebar(), router = useRouter()
const openMenu = ref(0)
const sidebarClass = computed(() =>
    'fixed top-0 left-0 z-20 flex flex-col flex-shrink-0 w-64 h-full pt-16 font-normal duration-75 lg:flex transition-width ' + (!sidebar.open ? 'hidden' : ''))
const menus = computed(() => router.menus.map(m => mapMenus(m)))

function mapMenus(m) {
    const r = { ...m }
    r.selected = router.path === r.link
    if (r.subs) {
        r.subs = r.subs.map(m2 => mapMenus(m2))
        r.selected = r.subs.some(m2 => m2.selected)
    }
    return r
}
</script>
<template>
    <aside :class="sidebarClass">
        <div
            class="relative flex flex-col flex-1 min-h-0 pt-0 bg-white border-r border-gray-200 dark:bg-gray-800 dark:border-gray-700">
            <div class="flex flex-col flex-1 pt-5 pb-4 overflow-y-auto">
                <div
                    class="flex-1 px-3 space-y-1 bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700">
                    <ul class="pb-2 space-y-2">
                        <SidebarMenu v-for="menu, i in menus" :key="i" :metadata="menu" :index="i" :open="i === openMenu"
                            @opening="v => { if (openMenu === v) openMenu = 0; else openMenu = v; }" />
                    </ul>
                </div>
            </div>
        </div>
    </aside>
</template>