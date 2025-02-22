import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDarkMode = defineStore('darkmode', () => {
    const isDark = ref(localStorage.getItem('color-theme') === 'dark' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches))
    function reverseTheme() {
        const theme = localStorage.getItem('color-theme')
        localStorage.setItem('color-theme', theme === 'dark' ? 'light' : 'dark')
        isDark.value = localStorage.getItem('color-theme') === 'dark' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
        if (theme === 'dark') document.documentElement.classList.remove('dark')
        else document.documentElement.classList.add('dark')
    }
    return { isDark, reverseTheme }
})
