import { defineConfig } from 'vite'
import path from 'path';
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'

import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
// https://vitejs.dev/config/
export default defineConfig({
  server: {
    host: '0.0.0.0'
  },
  plugins: [
    vue({
      reactivityTransform: [/src/]
    }),
    Components({
      resolvers: [NaiveUiResolver()]
    })
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
})
