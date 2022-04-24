import preprocess from 'svelte-preprocess'
import { resolve } from 'path'
// import adapter from '@sveltejs/adapter-auto'
import adapter from '@sveltejs/adapter-node'

let routeFolder = process.env.ROUTE_FOLDER

if (routeFolder == undefined) {
  routeFolder = 'guest'
}

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://github.com/sveltejs/svelte-preprocess
  // for more information about preprocessors
  preprocess: preprocess(),

  kit: {
    adapter: adapter({ out: '../.build/sveltekit/build' }),

    // Override http methods in the Todo forms
    methodOverride: {
      allowed: ['PUT', 'PATCH', 'DELETE'],
    },

    files: {
      routes: `src/routes/${routeFolder}`,
    },

    vite: {
      server: {
	      hmr: {
          port: 8081,
          clientPort: 8331,
        }
      },
      resolve: {
        alias: {
          $src: resolve('./src'),
          $lib: resolve('./src/lib'),
          $stores: resolve('./src/stores'),
          $assets: resolve('./src/assets'),
          $icon: resolve('./node_modules/svelte-bootstrap-icons/lib'),
        },
      },
    },
  },
}

export default config