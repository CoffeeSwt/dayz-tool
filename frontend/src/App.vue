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

</style>
