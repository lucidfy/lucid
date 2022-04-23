#!/bin/bash
(
    cd resources/views/svelte/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && npm audit fix \
    && npm run build \
) && \
cp -Rf resources/views/svelte/node_modules .build/sveltekit/
cp -Rf resources/views/svelte/package.json .build/sveltekit/
