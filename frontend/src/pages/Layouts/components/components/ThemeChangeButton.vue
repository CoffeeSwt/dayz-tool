<script setup lang="ts">
import { useThemeStore } from '@/store/themeStore';
const themeStore = useThemeStore()
themeStore.init()

</script>

<template>
    <div cursor-pointer>
        <div v-if="themeStore.theme == 'dark'" hover:bg-dark-hover size-7 flex-center rounded-lg
            @click="themeStore.changeTheme">
            <div size-5 bg-dark-icon class="i-line-md-sunny"></div>
        </div>
        <div v-if="themeStore.theme == 'light'" hover:bg-light-hover size-7 flex-center rounded-lg
            @click="themeStore.changeTheme">
            <div size-5 bg-light-icon class="i-solar-moon-line-duotone"></div>
        </div>
    </div>
</template>