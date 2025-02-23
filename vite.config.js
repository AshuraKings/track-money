import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    rollupOptions: {
      output: {
        manualChunks: id => {
          if (id.includes('node_modules')) {
            console.log(id)
            for (const s of ['pinia', 'oh-vue-icons', 'flowbite', 'apexcharts', 'datatables']) if (id.includes('/node_modules/' + s)) {
              return s
            }
            for (const s of [' vue']) if (id.includes('/node_modules/@' + s)) {
              return s
            }
            return 'vendor'
          }
        }
      }
    }
  }
})
