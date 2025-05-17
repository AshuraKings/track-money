import { createApp } from 'vue'
import './style.css'
import '../node_modules/flowbite-vue/dist/index.css'
import 'vue-loading-overlay/dist/css/index.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import './icons'
import { OhVueIcon } from "oh-vue-icons"
import VueApexCharts from 'vue3-apexcharts'

createApp(App)
    .use(createPinia())
    .use(VueApexCharts)
    .component('v-icon', OhVueIcon)
    .mount('#app')
