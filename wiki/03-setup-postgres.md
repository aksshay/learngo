# Postgres

Installed postgres for persisting user data.

## Database

host     = "192.168.1.10"
port     = 5432
user     = "postgres"
password = "password"
dbname   = "learngo"

## Queries

```SQL
CREATE DATABASE learngo OWNER postgres;
CREATE TABLE IF NOT EXISTS users  (
    id          integer NOT NULL PRIMARY KEY,
    user_name   varchar(25) NOT NULL,
    first_name  varchar(20) NOT NULL,
    last_name   varchar(20),
    password    varchar(50),
    email       varchar(60)
    created_at  timestamp DEFAULT current_timestamp,
    updated_at  timestamp DEFAULT current_timestamp,
    deleted_at  timestamp DEFAULT current_timestamp
);
```
