#!/bin/bash

cd /src/helia/pgmigrate
go run github.com/peterldowns/pgmigrate/cmd/pgmigrate@latest plan
