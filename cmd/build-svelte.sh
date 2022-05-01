#!/bin/bash
echo "Setting up svelte-kit!"
(
    cd svelte/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && LUCID_DOCKER_SETTING_UP=1 npm run build \
)

echo "Copying modules under .build/svelte-kit/"
rm -rf .build/svelte-kit/
cp -Rf svelte/node_modules .build/svelte-kit/
cp -Rf svelte/package.json .build/svelte-kit/

echo "Finished!"
