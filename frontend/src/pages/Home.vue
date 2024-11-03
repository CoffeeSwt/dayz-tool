<script setup lang="ts">
import Layouts from '@/pages/Layouts/index.vue'
import { ChooseFile, ChooseFileFolder } from '../../wailsjs/go/service/PathChooser'
import { useServerStartConfigStore } from '@/store/serverStartConfigStore';
const serverStartConfigStore = useServerStartConfigStore()
const ChooseServerPath = async () => {
    const res = await ChooseFileFolder()
    if (res == "") return
    console.log(res, serverStartConfigStore.config)
}
</script>

<template>
    <Layouts>
        <div @click="ChooseFile">ChooseFile</div>
        <div @click="ChooseServerPath">ChooseFileFolder</div>
    </Layouts>
</template>

<style scoped></style>
