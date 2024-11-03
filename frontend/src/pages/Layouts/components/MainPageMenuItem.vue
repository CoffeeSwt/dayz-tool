<script setup lang="ts">
import { computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

const props = defineProps({
    label: String,
    routerPath: String,
    icon: String,
})

const isActive = computed(() => {
    return router.currentRoute.value.fullPath == props.routerPath
})

const routeTo = () => {
    if (isActive.value) return
    router.push(props.routerPath!)
}

</script>

<template>
    <div w-full h-9 flex items-center p-2 rounded-lg cursor-pointer
        :class="{ 'bg-light-background-2   hover:bg-light-background-2 text-light-background-1 dark:bg-dark-background-2 dark:text-dark-text-background-2 dark:hover:bg-dark-background-2': isActive }"
        duration-100 @click="routeTo">
        <div size-6 mr-2 :class="props.icon"></div>
        <div>{{ props.label }}</div>
    </div>
</template>

<style scoped></style>
