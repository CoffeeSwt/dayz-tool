<script setup lang="ts">
import ButtonGroup from './components/ButtonGroup.vue';
import { useThemeStore } from '@/store/themeStore';

const themeStore = useThemeStore()
</script>

<template>
    <div size-full pr-8 flex justify-end dragarea items-center>
        <div h-full w-50 flex-center bg-light-background-1 dark:bg-dark-background-1>
            <div v-if="themeStore.theme == 'dark'" flex>
                <img h-12 aspect-2000-1284 src="../../assets/DayZ-logo-white.png" alt="">
                <div>TOOL</div>
            </div>
            <div v-if="themeStore.theme == 'light'" flex>
                <img h-12 aspect-3840-2160 src="../../assets/DayZ-logo-black.png" alt="">
                <div>TOOL</div>
            </div>
        </div>
        <div flex-grow-1></div>
        <div flex-shrink-0>
            <ButtonGroup></ButtonGroup>
        </div>
    </div>
</template>

<style scoped></style>
