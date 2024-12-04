<script setup lang="ts">
import { useThemeStore } from '@/store/themeStore';
import { useServerStartConfigStore } from '@/store/serverStartConfigStore.ts';
import { css } from '@/assets/fonts/SourceHanSansSC-Regular-2.otf'
import { onMounted } from 'vue';
const themeStore = useThemeStore()
const serverStartConfigStore = useServerStartConfigStore()
themeStore.init()
serverStartConfigStore.init()

onMounted(() => {
  //设置全局字体
  document.documentElement.style.fontFamily = `"${css.family}"`
})
</script>


<template>
  <RouterView></RouterView>
</template>