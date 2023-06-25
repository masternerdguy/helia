#!/bin/bash

export PGPASSWORD="$POSTGRES_PASSWORD"
echo 'drop schema public cascade;' | psql -h 127.0.0.1 -d $POSTGRES_DB -U $POSTGRES_USER 
echo 'create schema public;' | psql -h 127.0.0.1 -d $POSTGRES_DB -U $POSTGRES_USER 
psql -h 127.0.0.1 -d $POSTGRES_DB -U $POSTGRES_USER < /src/helia/devdb.sql
