<script setup lang="ts">
import ButtonGroup from './components/ButtonGroup.vue';
import { useThemeStore } from '@/store/themeStore';

const themeStore = useThemeStore()
</script>

<template>
    <div size-full pr-8 flex justify-end dragarea items-center>
        <div h-full w-50 flex-center bg-light-background-1 dark:bg-dark-background-1>
            <div v-if="themeStore.theme == 'dark'" flex items-center gap-2>
                <img h-12 aspect-3840-2160 src="@/assets/imgs/logo-white.png" alt="">
                <div text-size-4xl text-main>TOOL</div>
            </div>
            <div v-if="themeStore.theme == 'light'" flex items-center gap-2>
                <img h-12 aspect-3840-2160 src="@/assets/imgs/logo-black.png" alt="">
                <div text-size-4xl text-main>TOOL</div>
            </div>
        </div>
        <div flex-grow-1></div>
        <div flex-shrink-0>
            <ButtonGroup></ButtonGroup>
        </div>
    </div>
</template>

<style scoped></style>
