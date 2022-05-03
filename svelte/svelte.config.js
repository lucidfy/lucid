import preprocess from 'svelte-preprocess'
import { resolve } from 'path'
// import adapter from '@sveltejs/adapter-auto'
import adapter from '@sveltejs/adapter-node'

let envPort = 3000
let routeFolder = process.env.ROUTE_FOLDER

if (process.env.npm_config_argv !== undefined) {
  const argv = JSON.parse(process.env.npm_config_argv)
  envPort = parseInt(argv['remain'][parseInt(argv['remain'].indexOf('--port') + 1)])
}

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://github.com/sveltejs/svelte-preprocess
  // for more information about preprocessors
  preprocess: preprocess(),

  kit: {
    adapter: adapter({ out: `./build/${routeFolder}` }),

    // Override http methods in the Todo forms
    methodOverride: {
      allowed: ['PUT', 'PATCH', 'DELETE'],
    },

    files: {
      routes: `src/routes/${routeFolder}`,
      hooks: `src/hooks/lucid.ts`,
    },

    vite: {
      server: {
	      hmr: {
          port: envPort, // e.g: 8081
          clientPort: envPort + 250, // e.g 8081 + 250 = 8331
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
