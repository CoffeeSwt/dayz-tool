<script setup lang="ts">
import Layouts from '@/pages/Layouts/index.vue'
import { useColorMode } from '@vueuse/core';
const theme = useColorMode()



</script>

<template>
    <Layouts>
        <div size-20 class="bg-primary text-primary-foreground">{{ theme }}</div>
        <div size-20 class="bg-primary text-primary-foreground">primary</div>
        <div size-20 class="bg-secondary text-secondary-foreground">secondary</div>
        <div size-20 class="bg-accent text-accent-foreground">accent</div>
        <div size-20 class="bg-destructive text-destructive-foreground">destructive</div>
        <div size-20 class="bg-muted text-muted-foreground">muted</div>
        <div size-20 class="bg-popover text-popover-foreground">popover</div>
        <div size-20 class="bg-card text-card-foreground">popover</div>
    </Layouts>
</template>

<style scoped></style>
