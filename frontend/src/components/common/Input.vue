<template>
    <input :type="type" :value="modelValue" @input="onInput" :placeholder="placeholder" :disabled="disabled"
        :readonly="readonly" :autofocus="autofocus" outline-none dark:bg-dark-background-1 bg-light-background-1 p-1
        px-2 box-border border-gray dark:border-dark-background-1  border-solid border-1 rounded text-sm flex items-center />
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, PropType } from 'vue';

// 声明组件的 props 和事件
defineProps({
    modelValue: {
        type: [String, Number] as PropType<string | number>,
        default: ''
    },
    type: {
        type: String as PropType<'text' | 'password' | 'email' | 'number' | 'tel' | 'search'>,
        default: 'text'
    },
    placeholder: {
        type: String,
        default: ''
    },
    disabled: {
        type: Boolean,
        default: false
    },
    readonly: {
        type: Boolean,
        default: false
    },
    autofocus: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['update:modelValue']); // 声明事件

// 处理输入事件
const onInput = (event: Event) => {
    const input = event.target as HTMLInputElement;
    emit('update:modelValue', input.value);
};
</script>