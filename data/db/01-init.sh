#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS event (
	    id CHAR(26) NOT NULL CHECK (CHAR_LENGTH(id) = 26) PRIMARY KEY,
	    aggregate_id CHAR(26) NOT NULL CHECK (CHAR_LENGTH(aggregate_id) = 26),
	    event_data JSON NOT NULL,
	    version INT,
	    UNIQUE(aggregate_id, version)
	);
	CREATE INDEX idx_event_aggregate_id ON event (aggregate_id);
    CREATE TABLE IF NOT EXISTS users  (
        id          integer NOT NULL PRIMARY KEY,
        user_name   varchar(25) NOT NULL,
        first_name  varchar(20) NOT NULL,
        last_name   varchar(20),
        password    varchar(50),
        email       varchar(60),
        created_at  timestamp DEFAULT current_timestamp,
        updated_at  timestamp DEFAULT current_timestamp,
        deleted_at  timestamp DEFAULT current_timestamp
    );
  COMMIT;
EOSQL