import path from 'node:path'
import vue from '@vitejs/plugin-vue'
import autoprefixer from 'autoprefixer'

import tailwind from 'tailwindcss'
import { defineConfig } from 'vite'
import UnoCSS from 'unocss/vite'

import Font from 'vite-plugin-font';

export default defineConfig({
  css: {
    postcss: {
      plugins: [tailwind(), autoprefixer()],
    },
  },
  plugins: [
    vue(),
    UnoCSS(),
    Font.vite({
      scanFiles: ['src/**/*.{vue,ts,tsx,js,jsx}'],
    })
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})