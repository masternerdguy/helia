#!/bin/bash

# login to registry
docker login heliaalpharegistry.azurecr.io --username 'heliaalpharegistry' --password 'u441Vgr2AEasyW0fLqCSQcdJPL5XQ=pH'

# tag local image
docker tag heliaalpha heliaalpharegistry.azurecr.io/heliaalpha

# push local image to azure
docker push heliaalpharegistry.azurecr.io/heliaalpha

# all done!
echo 'GAME ENGINE DEPLOYMENT COMPLETE :party parrot:'
