import path from 'path';
import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter(),
		alias: {
			$lib: path.resolve('./src/lib'),
			$data: path.resolve('./src/data'),
			$hooks: path.resolve('./src/hooks'),
			$types: path.resolve('./src/types'),
			$stores: path.resolve('./src/stores'),
			$components: path.resolve('./src/components')
		}
	}
};

export default config;
