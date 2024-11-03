<script setup lang="ts">
import { useThemeStore } from '@/store/themeStore';
import { useServerStartConfigStore } from '@/store/serverStartConfigStore.ts';
const themeStore = useThemeStore()
const serverStartConfigStore = useServerStartConfigStore()
themeStore.init()
serverStartConfigStore.init()
</script>


<template>
  <RouterView></RouterView>
</template>