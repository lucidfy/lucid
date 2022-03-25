#!/bin/bash
NEEDED_COMMANDS="go npm node"

for cmd in ${NEEDED_COMMANDS} ; do
    if ! command -v ${cmd} &> /dev/null ; then
        echo Please install ${cmd}!
        exit 1
    fi
done
