import {
    defineConfig,
    presetAttributify,
    presetIcons,
    presetTypography,
    presetUno,
    presetWebFonts,
    transformerDirectives,
    transformerVariantGroup
} from 'unocss'

export default defineConfig({
    rules: [
        // ['m-1', { margin: '0.25rem' }],
        // [/^m-(\d+)$/, ([, d]) => ({ margin: `${d / 4}rem` })],
        // [/^p-(\d+)$/, match => ({ padding: `${match[1] / 4}rem` })],
        ['dragarea', { '--wails-draggable': 'drag' }],
    ],
    shortcuts: [
        ['flex-center', ['flex justify-center items-center']]
    ],
    theme: {
        colors: {
            'light-icon': '#030712',
            'light-hover': '#f3f4f6',
            'dark-icon': '#f9fafb',
            'dark-hover': '#1f2937'
        }
    },
    presets: [
        presetUno(),
        presetAttributify(),
        presetIcons(),
        presetTypography(),
        presetWebFonts({
            fonts: {
                // ...
            },
        }),
    ],
    transformers: [
        transformerDirectives(),
        transformerVariantGroup(),
    ],
})