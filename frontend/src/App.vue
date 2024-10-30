<script setup lang="ts">
import { useThemeStore } from '@/store/themeStore';
const themeStore = useThemeStore()
themeStore.init()
</script>


<template>
  <RouterView></RouterView>
</template>