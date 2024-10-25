<script setup lang="ts">
import { WindowMinimise } from '../../../../../wailsjs/runtime'
const minimise = () => {
    WindowMinimise()
}
</script>

<template>
    <div cursor-pointer hover:bg-light-hover dark:hover:bg-dark-hover size-7 flex-center rounded-lg @click="minimise">
        <div size-5 dark:bg-dark-icon bg-light-icon class="i-ic-baseline-minus"></div>
    </div>
</template>