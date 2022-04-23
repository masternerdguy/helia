#!/bin/bash

# login to registry
docker login heliaalpharegistry.azurecr.io --username 'heliaalpharegistry' --password 'wN/oE7R0+3jxiCKAxfIcdEWv4tqHwUu1'

# tag local image
docker tag heliaalpha heliaalpharegistry.azurecr.io/heliaalpha

# push local image to azure
docker push heliaalpharegistry.azurecr.io/heliaalpha

# inform user to update startup timeout
echo 'GAME ENGINE DEPLOYMENT COMPLETE :party parrot:'
echo 'Please run the following on the Azure CLI to avoid a startup timeout: '
echo '    az webapp config appsettings set --resource-group helia --name helia-alpha-engine --settings WEBSITES_CONTAINER_START_TIME_LIMIT=1800'
