#!/bin/bash

# get to repository root
cd ~/go/src/helia

# clean repository
git stash
git clean -fxd

# copy alpha config
cp -v ~/go/src/helia/db-configuration-alpha.json ~/go/src/helia/db-configuration.json

# remove frontend and scripting files
rm -rfv ~/go/src/helia/frontend
rm -rfv ~/go/src/helia/bash

# remove sql backups
rm -rfv ~/go/src/helia/*.sql

# build docker image
docker build -t heliaalpha . -f Dockerfile.single

# restore repo state
git clean -fxd
git reset --hard
git stash pop
