#!/bin/bash

export PGPASSWORD='thisiswherewefight' 
echo 'drop owned by developer;' | psql -h 127.0.0.1 -d helia -U developer 
psql -h 127.0.0.1 -d helia -U developer < ../devdb.sql
