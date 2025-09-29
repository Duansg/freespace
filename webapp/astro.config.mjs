import { defineConfig } from 'astro/config';
import preact from '@astrojs/preact';
import tailwind from '@astrojs/tailwind';
import node from '@astrojs/node';

// Determine the API target based on environment
const getApiTarget = () => {
  // In Docker environment, use service name
  if (process.env.NODE_ENV === 'production') {
    return 'http://backend:8888';
  }
  // In development, use localhost
  return 'http://localhost:8888';
};

// https://astro.build/config
export default defineConfig({
	integrations: [preact(), tailwind()],
	output: 'server',
	adapter: node({ mode: 'standalone' }),
	server: {
		host: '0.0.0.0',
		port: 4321,
	},
	preview: {
		host: '0.0.0.0',
		port: 4321,
	},
	vite: {
		server: {
			proxy: {
				'/api': {
					target: getApiTarget(),
					changeOrigin: true
				}
			}
		}
	}
});
