#!/bin/bash
(
    cd resources/sveltekit-starter/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && npm audit fix \
    && npm run build \
) && \
cp -Rf resources/sveltekit-starter/node_modules .build/sveltekit/
