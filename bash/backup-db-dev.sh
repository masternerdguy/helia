#!/bin/bash

export PGPASSWORD='fdb01db4749748cfbdb21e4766570561!'
pg_dump -h 127.0.0.1 -d helia -U heliaagent -f ../devschema.sql -s
pg_dump -h 127.0.0.1 -d helia -U heliaagent -f ../devdb.sql --column-inserts --data-only

# --column-inserts --data-only
