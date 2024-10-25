<script setup lang="ts">
import { Quit } from '../../../../../wailsjs/runtime'
const closeAPP = () => {
    Quit()
}
</script>

<template>
    <div cursor-pointer hover:bg-light-hover dark:hover:bg-dark-hover size-7 flex-center rounded-lg @click="closeAPP">
        <div size-5 dark:bg-dark-icon bg-light-icon class="i-line-md-close"></div>
    </div>
</template>