<script setup lang="ts">
import ThemeChangeButton from './components/ThemeChangeButton.vue'
import MaximizeButton from './components/ToggleMaximizeButton.vue'
import MinimizeButton from './components/MinimizeButton.vue'
import CloseButton from './components/CloseButton.vue';
</script>

<template>
    <div size-full flex items-center>
        <ThemeChangeButton></ThemeChangeButton>
        <MinimizeButton></MinimizeButton>
        <MaximizeButton></MaximizeButton>
        <CloseButton></CloseButton>
    </div>
</template>

<style scoped></style>
