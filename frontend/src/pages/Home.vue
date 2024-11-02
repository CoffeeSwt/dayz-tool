<script setup lang="ts">
import Layouts from '@/pages/Layouts/index.vue'
import { ChooseFile, ChooseFileFolder } from '../../wailsjs/go/service/PathChooser'
</script>

<template>
    <Layouts>
        <div @click="ChooseFile">ChooseFile</div>
        <div @click="ChooseFileFolder">ChooseFileFolder</div>
    </Layouts>
</template>

<style scoped></style>
