import { useColorMode } from '@vueuse/core'

export function changeTheme(e: MouseEvent) {
    // Pass { disableTransition: false } to enable transitions
    const currentTheme = useColorMode()
    const root = document.documentElement;
    // // 使用setProperty设置CSS变量的值
    root.style.setProperty('--x', `${e.clientX}px`);
    root.style.setProperty('--y', `${e.clientY}px`);

    const change = () => {
        if (currentTheme.value == 'light') {
            currentTheme.value = 'dark'
        } else {
            currentTheme.value = 'light'
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