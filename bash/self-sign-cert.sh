#!/bin/bash

# generate self signed cert based on config
openssl req -config ~/go/src/helia/ssl-cert.conf -new -x509 -newkey rsa:2048 -nodes -keyout ~/go/src/helia/ssl.key.pem -days 365 -out ~/go/src/helia/ssl.cert.pem

# all done!
echo "Eyup..."
