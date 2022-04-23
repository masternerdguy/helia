#!/bin/bash

export PGPASSWORD='thisiswherewefight'
pg_dump -h 127.0.0.1 -d helia -U developer -f ../devschema.sql -s
pg_dump -h 127.0.0.1 -d helia -U developer -f ../devdb.sql
