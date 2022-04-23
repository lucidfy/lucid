#!/bin/bash
(
    cd .svelte-kit/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && npm audit fix \
    && npm run build \
) && \
cp -Rf .svelte-kit/node_modules .build/sveltekit/
cp -Rf .svelte-kit/package.json .build/sveltekit/
