<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Greet } from '../wailsjs/go/main/App'
import { ref } from 'vue'
const data = ref('')
Greet('hello').then(res => {
  data.value = res
})
</script>

<template>
  <Button>111</Button> 
  <div>{{ data }}</div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
