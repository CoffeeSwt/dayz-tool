import { useColorMode } from "@vueuse/core"
import { defineStore } from "pinia"
import { ref } from "vue"

export const useThemeStore = defineStore('theme', () => {
    const theme = ref('light')
    function init() {
        const res = localStorage.getItem('vueuse-color-scheme')
        if (res == null || res == 'auto') {
            const mode = useColorMode()
            mode.value = 'light'
        } else {
            theme.value = res
        }
    }
    function changeTheme(e: MouseEvent) {
        // Pass { disableTransition: false } to enable transitions
        const currentTheme = useColorMode()
        const root = document.documentElement;
        // // 使用setProperty设置CSS变量的值
        root.style.setProperty('--x', `${e.clientX}px`);
        root.style.setProperty('--y', `${e.clientY}px`);

        const change = () => {
            if (currentTheme.value == 'light') {
                currentTheme.value = 'dark'
                theme.value = 'dark'
            } else {
                currentTheme.value = 'light'
                theme.value = 'light'
            }
        }
        //@ts-ignore
        //ViewTransition API兼容性检查
        if (!document.startViewTransition) {
            change()
        } else {
            //@ts-ignore
            document.startViewTransition(() => {
                change()
            })
        }
    }

    return { theme, init, changeTheme }
})