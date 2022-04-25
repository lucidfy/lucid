#!/bin/bash
echo "Setting up svelte-kit!"
(
    cd svelte/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && npm run build \
)

echo "Copying modules under .build/svelte-kit/"
rm -rf .build/svelte-kit/
cp -Rf svelte/node_modules .build/svelte-kit/
cp -Rf svelte/package.json .build/svelte-kit/

echo "Finished!"
