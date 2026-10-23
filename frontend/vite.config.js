import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 4582,
    strictPort: true,
    proxy: {
      '/api': process.env.API_URL ?? 'http://localhost:4581',
    },
  },
})
