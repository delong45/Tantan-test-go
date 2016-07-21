#/bin/bash

source env.sh
psql -U postgres -f sql/user.sql
go build tantan
