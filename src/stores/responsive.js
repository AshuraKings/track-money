import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useResponsive = defineStore('responsive', () => {
    const windowWidth = ref(window.innerWidth)
    function setWidth(width) {
        windowWidth.value = width
    }
    const isLargeScreen = computed(() => windowWidth.value >= 1200)
    const isMediumScreen = computed(() => windowWidth.value >= 768 && windowWidth.value < 1200)
    const isSmallScreen = computed(() => windowWidth.value < 768)
    return { windowWidth, setWidth, isLargeScreen, isMediumScreen, isSmallScreen }
})
