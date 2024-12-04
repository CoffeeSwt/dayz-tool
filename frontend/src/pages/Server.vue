<script setup lang="ts">
import Layouts from '@/pages/Layouts/index.vue'
import Input from '@/components/common/Input.vue';
import { ChooseFileFolder } from '../../wailsjs/go/service/PathChooser'
import { ref } from 'vue';
import {
    Tooltip,
    TooltipContent,
    TooltipProvider,
    TooltipTrigger,
} from '@/components/ui/tooltip'

const dayzserverpath = ref('')

const chooseFileFolderHandler = async () => {
    dayzserverpath.value = await ChooseFileFolder()
}

const clearForm = () => {
    dayzserverpath.value = ''
}

</script>

<template>
    <Layouts>
        <div>
            <div flex items-center gap-2>
                <div>服务器根目录</div>
                <Input :model-value="dayzserverpath"></Input>
                <TooltipProvider :delayDuration="500">
                    <Tooltip>
                        <TooltipTrigger as-child>
                            <div @click="chooseFileFolderHandler" class="group/item" flex-center size-8 border-rounded
                                hover:bg-main duration-300 cursor-pointer>
                                <div size-6 i-material-symbols-folder-open duration-300
                                    class="group-hover/item:bg-dark-text-background-2">
                                </div>
                            </div>
                        </TooltipTrigger>
                        <TooltipContent dark:border-dark-background>
                            <p>选择目录</p>
                        </TooltipContent>
                    </Tooltip>
                </TooltipProvider>
                <TooltipProvider :delayDuration="500">
                    <Tooltip>
                        <TooltipTrigger as-child>
                            <div @click="clearForm" class="group/item" flex-center size-8 border-rounded
                                hover:bg-red duration-300 cursor-pointer>
                                <div size-6 i-material-symbols-cleaning-services-rounded duration-300
                                    class="group-hover/item:bg-dark-text-background-2">
                                </div>
                            </div>
                        </TooltipTrigger>
                        <TooltipContent dark:border-dark-background>
                            <p>清除</p>
                        </TooltipContent>
                    </Tooltip>
                </TooltipProvider>
            </div>
        </div>
    </Layouts>
</template>

<style scoped></style>
