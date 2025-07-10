import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'node:path'

import tailwindcss from '@tailwindcss/vite';

const src_path = path.resolve(__dirname, './src');

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@": src_path
    }
  },
  build: {
    // generate .vite/manifest.json in outDir
    manifest: true,
    rollupOptions: {
      // overwrite default .html entry
      input: './src/main.js',
    },
  },
  server: {
    origin: 'http://localhost:8080',
  },
  plugins: [vue(), tailwindcss()],
})
