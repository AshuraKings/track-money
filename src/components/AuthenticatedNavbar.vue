<script setup>
import DarkModeNav from './DarkModeNav.vue';
import NavStart from './NavStart.vue';
import ProfileNav from './ProfileNav.vue';

</script>
<template>
    <nav class="fixed z-30 w-full bg-white border-b border-gray-200 dark:bg-gray-800 dark:border-gray-700">
        <div class="px-3 py-3 lg:px-5 lg:pl-3">
            <div class="flex items-center justify-between">
                <NavStart />
                <div class="flex items-center">
                    <div class="hidden mr-3 -mb-1 sm:block">
                        <span>
                            <!-- NavMode -->
                        </span>
                    </div>
                    <DarkModeNav />
                    <ProfileNav />
                </div>
            </div>
        </div>
    </nav>
</template>