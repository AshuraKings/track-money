import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    rollupOptions: {
      output: {
        manualChunks: id => {
          console.log(id)
          if (id.includes('node_modules')) {
            for (const s of ['pinia', 'oh-vue-icons', 'flowbite', 'apexcharts', 'datatables', 'vue-loading-overlay']) if (id.includes('/node_modules/' + s)) {
              return s
            }
            for (const s of [' vue']) if (id.includes('/node_modules/@' + s)) {
              return s
            }
            return 'vendor'
          }
          if (id.endsWith('.vue')) {
            const dirs = id.split('/')
            const parent = dirs[dirs.length - 2]
            const s = dirs[dirs.length - 1].replace('.vue', '')
            return parent + '/' + s
          }
        }
      }
    }
  }
})
