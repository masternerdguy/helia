#!/bin/bash

export PGPASSWORD='fdb01db4749748cfbdb21e4766570561!' 
export PGSSLMODE='require' 

pg_dump -h helia-alpha.postgres.database.azure.com -d helia -U heliaagent@helia-alpha -f ../alphaschema.sql -s 
pg_dump -h helia-alpha.postgres.database.azure.com -d helia -U heliaagent@helia-alpha -f ../alphadb.sql
