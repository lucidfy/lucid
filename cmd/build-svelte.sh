#!/bin/bash
(
    cd svelte/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && npm audit fix \
    && npm run build \
) && \
cp -Rf svelte/node_modules .build/svelte-kit/
cp -Rf svelte/package.json .build/svelte-kit/
