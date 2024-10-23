import path from 'node:path'
import vue from '@vitejs/plugin-vue'
import autoprefixer from 'autoprefixer'

import tailwind from 'tailwindcss'
import { defineConfig } from 'vite'
import UnoCSS from 'unocss/vite'

export default defineConfig({
  css: {
    postcss: {
      plugins: [tailwind(), autoprefixer()],
    },
  },
  plugins: [vue(), UnoCSS(),],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})