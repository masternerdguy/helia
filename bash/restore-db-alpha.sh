#!/bin/bash

export PGPASSWORD='fdb01db4749748cfbdb21e4766570561!' 
export PGSSLMODE='require' 
echo 'drop owned by heliaagent;' | psql -h helia-alpha.postgres.database.azure.com -d helia -U heliaagent@helia-alpha 
psql -h helia-alpha.postgres.database.azure.com -d helia -U heliaagent@helia-alpha < ../alphadb.sql
