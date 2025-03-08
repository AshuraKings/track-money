<template>
    <li :class="metadata.selected ? 'bg-smoke dark:bg-gray-700 rounded-lg' : ''">
        <button type="button" v-if="metadata.subs" @click.prevent="() => $emit('opening', index)"
            class="flex items-center w-full p-2 text-base text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-gray-700">
            <v-icon v-if="metadata.icon" :name="metadata.icon"
                class="w-6 h-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white" />
            <span :class="textClass" v-text="metadata.label"></span>
            <svg sidebar-toggle-item="" class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                    d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                    clip-rule="evenodd"></path>
            </svg>
        </button>
        <ul v-if="metadata.subs" :class="[ulClass]">
            <li v-for="sub, i in metadata.subs" :key="i"
                :class="sub.selected ? 'bg-smoke dark:bg-gray-700 rounded-lg' : ''">
                <Link :href="sub.href"
                    :class="'text-base text-gray-900 rounded-lg flex items-center p-2 group hover:bg-gray-100 transition duration-75 '
                        + 'pl-11 dark:text-gray-200 dark:hover:bg-gray-700' + (sub.selected ? 'bg-smoke dark:bg-gray-700' : '')">
                {{ sub.label }}
                </Link>
            </li>
        </ul>
        <Link v-else :href="metadata.link"
            class="flex items-center p-2 text-base text-gray-900 rounded-lg hover:bg-gray-100 group dark:text-gray-200 dark:hover:bg-gray-700">
        <v-icon v-if="metadata.icon" :name="metadata.icon"
            class="w-6 h-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white" />
        <span :class="[textClass]" v-text="metadata.label" />
        </Link>
    </li>
</template>
<script setup>
import { computed } from 'vue'
import Link from './Link.vue'

const props = defineProps({
    metadata: Object,
    index: Number,
    open: Boolean
})

const textClass = computed(() => props.metadata.subs
    ? 'flex-1 ml-3 text-left whitespace-nowrap'
    : 'ml-3'
)
const ulClass = computed(() => 'space-y-2 py-2 ' + (props.open ? '' : 'hidden'))
</script>