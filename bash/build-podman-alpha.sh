#!/bin/bash

# build podman image
podman build -t heliaalpha . -f podmanfile.prodengine
