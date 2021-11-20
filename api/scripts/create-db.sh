#!/bin/bash
set -e

if [ ! -f "/docker-entrypoint-initdb.d/backup.sql" ]; then

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" --password "$POSTGRES_PASSWORD" -- <<-EOSQL
    CREATE TABLE barbers (
        id serial primary key,
	    last_name varchar(100) not null,
	    first_name varchar(100) not null,
	    age INTEGER not null,
	    description varchar(100) not null,
	    created timestamp without time zone default (now() at time zone 'utc')
    );
EOSQL
fi
