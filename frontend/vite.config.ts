import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
		// Make it possible to use '@' in import statements
        alias: {
            '@': path.resolve(__dirname, './src'),
        }
    },
    build: {
		// On final build, use hashes as the filenames
        rollupOptions: {
            output: {
                entryFileNames: `assets/[hash].js`,
                chunkFileNames: `assets/[hash].js`,
                assetFileNames: `assets/[hash].[ext]`
            }
    	}
    },
	server: {
		// Proxy developement server
		proxy: {
			'/api': 'http://localhost:8080',
            '/files': 'http://localhost:8080'
        }
	}
})
