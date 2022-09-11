#!/bin/bash

echo "this script is intended for use in dev/demo environments only!"

sleep 5

killall serve-all.sh
killall 'npm start'
killall ng
killall main

