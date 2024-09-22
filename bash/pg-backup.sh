#!/bin/bash

export PGPASSWORD="$POSTGRES_PASSWORD"

pg_dump -h 127.0.0.1 -d $POSTGRES_DB -U $POSTGRES_USER  -f /src/helia/devschema.sql -s
pg_dump -h 127.0.0.1 -d $POSTGRES_DB -U $POSTGRES_USER  -f /src/helia/devdb.sql #--column-inserts --data-only

# --column-inserts --data-only
