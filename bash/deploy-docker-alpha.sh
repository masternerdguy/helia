#!/bin/bash

# login to registry
docker login heliaalpharegistry.azurecr.io

# tag local image
docker tag heliaalpha heliaalpharegistry.azurecr.io/heliaalpha

# push local image to azure
docker push heliaalpha
