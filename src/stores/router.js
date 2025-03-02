import { defineStore } from "pinia"
import { ref } from "vue"

export const useRouter = defineStore('router', () => {
    const path = ref('/'), token = ref(localStorage.getItem('sessionToken') || ''), loading = ref(false), title = ref('App')
    function setPath(newPath = '') {
        path.value = newPath
    }
    function setToken(newToken = '', refreshToken = '') {
        if (newToken) {
            localStorage.setItem('sessionToken', newToken)
            localStorage.setItem('refreshToken', refreshToken)
        } else localStorage.clear()
        token.value = newToken
    }
    function reverseLoading() {
        loading.value = !loading.value
    }
    function setTitle(newTitle = '') {
        title.value = newTitle
        document.title = newTitle
    }
    return { loading, path, reverseLoading, setPath, setTitle, setToken, token }
})