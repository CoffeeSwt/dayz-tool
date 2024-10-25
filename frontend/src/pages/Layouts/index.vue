<script setup lang="ts">
import Header from './Header.vue';
import Side from './Side.vue';
</script>

<template>
    <div w-full h-screen flex flex-col overflow-hidden class="bg-background text-foreground">
        <div w-full h-20 flex-shrink-0>
            <Header></Header>
        </div>
        <div flex-grow-1 overflow-auto flex>
            <div w-50 h-full flex-shrink-0>
                <Side></Side>
            </div>
            <div flex-grow-1>
                <slot></slot>
            </div>
        </div>
    </div>

</template>

<style scoped></style>
