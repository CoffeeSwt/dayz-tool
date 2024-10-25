<script setup lang="ts">
import { WindowToggleMaximise } from '../../../../../wailsjs/runtime/runtime'
const toggleMaximise = () => {
    WindowToggleMaximise()
}
</script>

<template>
    <div cursor-pointer hover:bg-light-hover dark:hover:bg-dark-hover size-7 flex-center rounded-lg
        @click="toggleMaximise">
        <div size-5 dark:bg-dark-icon bg-light-icon class="i-material-symbols-chrome-maximize-outline"></div>
    </div>
</template>