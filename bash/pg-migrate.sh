#!/bin/bash

cd /src/helia/flyway
go run github.com/peterldowns/pgmigrate/cmd/pgmigrate@latest migrate
