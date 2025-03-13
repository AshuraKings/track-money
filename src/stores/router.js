import { defineStore } from "pinia"
import { ref } from "vue"

export const useRouter = defineStore('router', () => {
  const path = ref(localStorage.getItem('sessionToken') ? '/dashboard' : '/'), token = ref(localStorage.getItem('sessionToken') || ''), loading = ref(false), title = ref('')
  const username = ref('username'), name = ref('name'), role = ref('admin'), menus = ref([])
  function setPath(newPath = '') {
    path.value = newPath
  }
  function setToken(newToken = '', refreshToken = '') {
    if (newToken) {
      localStorage.setItem('sessionToken', newToken)
      localStorage.setItem('refreshToken', refreshToken)
    } else {
      localStorage.clear()
      username.value = name.value = role.value = ''
      menus.value = []
    }
    token.value = newToken
  }
  function reverseLoading() {
    loading.value = !loading.value
  }
  function setTitle(newTitle = '') {
    title.value = newTitle
    document.title = newTitle
  }
  function setSession(session = { id: 0, role: { id: 0, nm: '' }, user: { id: 0, nm: '', username: '' } }) {
    username.value = session.user.username
    name.value = session.user.nm
    role.value = session.role.nm
  }
  function isRedirectNotRequired(menu = {}) {
    return menu.link === path.value || (
      menu.subs &&
      menu.subs.some(m => isRedirectNotRequired(m))
    )
  }
  function setMenus(menu = []) {
    menus.value = menu
    if (!menu || !menu.some(m => isRedirectNotRequired(m))) path.value = '/dashboard'
  }
  return { loading, menus, name, path, reverseLoading, role, setMenus, setPath, setSession, setTitle, setToken, token, username }
})