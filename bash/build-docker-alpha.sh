#!/bin/bash

# get to repository root
cd ~/go/src/helia

# clean repository
git stash
git clean -fxd

# copy alpha config
cp -v ~/go/src/helia/db-configuration-alpha.json ~/go/src/helia/db-configuration.json
cp -v ~/go/src/helia/listener-configuration-alpha.json ~/go/src/helia/listener-configuration.json

# remove frontend and scripting files
rm -rfv ~/go/src/helia/frontend
rm -rfv ~/go/src/helia/bash

# temporarily move .git out of repo
mv ~/go/src/helia/.git ~/tmpgit

# remove sql backups
rm -rfv ~/go/src/helia/*.sql

# build docker image
docker build -t heliaalpha . -f Dockerfile.single

# move .git back into repo
mv ~/tmpgit ~/go/src/helia/.git

# restore repo state
git clean -fxd
git reset --hard

# restore executable permission
chmod a+x bash/build-docker-alpha.sh
