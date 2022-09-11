#!/bin/bash

echo "this script is intended for use in dev/demo environments only!"
sleep 5

# start frontend
echo "* starting angular site..."

cd ~/go/src/helia/frontend
yarn install
npm start &

# start backend
echo "* starting engine server..."

cd ..

while true
do
      go run main.go
done

