<script setup>
import { useRouter } from './stores/router'
import Login from './pages/Login.vue'
import Register from './pages/Register.vue'
import Loading from 'vue-loading-overlay'
import Dashboard from './pages/Dashboard.vue'
import Users from './pages/Users.vue'
import Roles from './pages/Roles.vue'
import Menus from './pages/Menus.vue'

const router = useRouter()
</script>
<template>
  <Loading v-if="router.loading" :can-cancel="false" is-full-page v-model:active="router.loading" />
  <Dashboard v-if="router.path === '/dashboard'" />
  <Register v-if="router.path === '/register'" />
  <Login v-if="router.path === '/'" />
  <Users v-if="router.path === '/master/users'" />
  <Roles v-if="router.path === '/master/roles'" />
  <Menus v-if="router.path === '/master/menus'" />
</template>