<script setup lang="ts">
import Header from './Header.vue';
import Side from './Side.vue';
</script>

<template>
    <div w-full h-screen flex>
        <div w-40 flex-shrink-0 h-full>
            <Side></Side>
        </div>
        <div flex-grow-1 h-40 flex flex-col>
            <Header></Header>
            <div>
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<style scoped></style>
