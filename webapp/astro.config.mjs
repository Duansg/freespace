import { defineConfig } from 'astro/config';
import preact from '@astrojs/preact';
import tailwind from '@astrojs/tailwind';
import node from '@astrojs/node';

// https://astro.build/config
export default defineConfig({
	integrations: [preact(), tailwind()],
	output: 'server',
	adapter: node({ mode: 'standalone' }),
	server: {
		host: '0.0.0.0',
		port: 8241,
	},
	preview: {
		host: '0.0.0.0',
		port: 8241,
	},
	vite: {
		server: {
			proxy: {
				'/api': {
					target: 'http://localhost:8240',
					changeOrigin: true
				}
			}
		}
	}
});
