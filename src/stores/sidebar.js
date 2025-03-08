import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSidebar = defineStore('sidebar', () => {
    const open = ref(false)
    function setOpenClose() {
        open.value = !open.value
    }
    return { open, setOpenClose }
})
