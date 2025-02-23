import { createApp } from 'vue'
import './style.css'
import '../node_modules/flowbite-vue/dist/index.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import './icons'
import { OhVueIcon } from "oh-vue-icons"

createApp(App)
    .use(createPinia())
    .component('v-icon', OhVueIcon)
    .mount('#app')
