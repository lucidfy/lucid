#!/bin/bash
echo "Firing up svelte-kit!"
(
    cd svelte/ \
    && rm -rf .svelte-kit/ \
    && npm install \
    && LUCID_DOCKER_SETTING_UP=1 ROUTE_FOLDER=guest npm run build \
    && LUCID_DOCKER_SETTING_UP=1 ROUTE_FOLDER=docs npm run build \
)
echo "Finished!"
