<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Greet } from '../wailsjs/go/main/App'
import { GetLocalConfig } from '../wailsjs/go/service/ConfigService'
import { ref } from 'vue'
const data = ref('')
const data1 = ref('')
Greet('hello').then(res => {
  data.value = res
})
const get = () => {
  GetLocalConfig().then(res => {
    console.log(res)
  })
}


</script>

<template>
  <Button @click="get">111</Button>
  <div>{{ data }}</div>
  <div>{{ data1 }}</div>
</template>

<style scoped></style>
