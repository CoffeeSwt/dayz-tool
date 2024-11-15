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
  document.body.style.fontFamily = `"${css.family}"`
})
</script>


<template>
  <RouterView></RouterView>
</template>