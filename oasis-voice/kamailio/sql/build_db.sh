#!/bin/sh

cat standard-create.sql permissions-create.sql carriers.sql | PGPASSWORD=$SQL_PASSWORD  psql -h $SQL_HOST $SQL_USER $SQL_DATABASE