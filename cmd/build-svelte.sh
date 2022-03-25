#!/bin/bash
(cd resources/sveltekit-starter/ && npm install && npm audit fix && npm run build)
