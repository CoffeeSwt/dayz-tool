<script setup lang="ts">
import MainPageMenuItem from './components/MainPageMenuItem.vue';
</script>

<template>
    <div size-full bg-light-background-1 dark:bg-dark-background-1 p-2 flex flex-col gap-2>
        <MainPageMenuItem router-path="/" label="服务器管理" icon="i-bx-bxs-server"></MainPageMenuItem>
        <MainPageMenuItem router-path="/map" label="在线地图" icon="i-ic-round-location-on"></MainPageMenuItem>
        <MainPageMenuItem router-path="/wiki" label="游戏百科" icon="i-material-symbols-menu-book"></MainPageMenuItem>
    </div>
</template>

<style scoped></style>
