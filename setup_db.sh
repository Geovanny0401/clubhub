#!/bin/bash
set -e

# Parámetros de conexión a PostgreSQL
export PGPASSWORD=postgres
HOST="clubhubdb"
USER="postgres"
DB="clubhub"
PORT="5432"

# Conectándose a PostgreSQL y ejecutando comandos SQL
psql -h $HOST -U $USER -p $PORT -d $DB <<-EOSQL
    CREATE DATABASE clubhub;
    GRANT ALL PRIVILEGES ON DATABASE clubhub TO clubhub;
EOSQL