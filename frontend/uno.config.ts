import { match } from 'assert'
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
        // ['aspect-2600-1284', { 'aspect-raito': '2600/1284' }],
        [/^aspect-(\d+)-(\d+)$/, match => ({ 'aspect-ratio': `${match[1]}/${match[2]}` })],
    ],
    shortcuts: [
        ['flex-center', ['flex justify-center items-center']]
    ],
    theme: {
        colors: {

            //light background
            'light-icon': '#030712',
            'light-hover': '#f3f4f6',
            'light-background': '#f7f9fc',
            'light-text': '#000000',
            //light-background-1
            'light-background-1': '#f0f3f6',
            'light-hover-background-1': '#e4e8ec',
            'light-text-background-1': '#737b89',


            //dark background
            'dark-background': '#13131a',
            'dark-text': '#ffffff',
            'dark-icon': '#f9fafb',
            'dark-hover': '#1f2937',
            //dark-background-1
            'dark-background-1': '#1a1a21',
            'dark-hover-background-1': '#27272e',
            'dark-text-background-1': '#a9a9ab',


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