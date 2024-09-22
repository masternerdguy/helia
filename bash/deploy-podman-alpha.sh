#!/bin/bash

# login to registry
podman login $HELIA_ALPHA_REGISTRY --username $HELIA_ALPHA_REGISTRY_USER --password $HELIA_ALPHA_REGISTRY_PASS

# tag local image
podman tag heliaalpha "$HELIA_ALPHA_REGISTRY/heliaalpha"

# push local image to azure
podman push "$HELIA_ALPHA_REGISTRY/heliaalpha"

# all done!
echo 'GAME ENGINE DEPLOYMENT COMPLETE :party parrot:'
