#!/bin/bash

killall serve-all.sh
killall 'npm start'
killall 'ng serve --host 0.0.0.0 --disable-host-check'
killall main

echo "this script is intended for use in dev/demo environments only!"

