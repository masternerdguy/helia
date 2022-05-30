#!/bin/bash

export PGPASSWORD='fdb01db4749748cfbdb21e4766570561!' 
echo 'drop owned by heliaagent;' | psql -h 127.0.0.1 -d helia -U heliaagent 
psql -h 127.0.0.1 -d helia -U heliaagent < ../devdb.sql
