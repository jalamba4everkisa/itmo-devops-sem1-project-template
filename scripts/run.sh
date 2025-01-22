#!/bin/bash
export PGPASSWORD=$POSTGRES_PASSWORD
systemctl start postgresql
psql -U $POSTGRES_USER -c 'CREATE DATABASE "'$POSTGRES_DB'";'
psql -U $POSTGRES_USER -d $POSTGRES_DB -h $POSTGRES_HOST -c "CREATE TABLE prices(
                                        id INTEGER PRIMARY KEY NOT NULL,
                                        name VARCHAR NOT NULL,
                                        category VARCHAR NOT NULL,
                                        price REAL NOT NULL,
                                        create_date DATE NOT NULL
                                    )";

psql -U $POSTGRES_USER -d $POSTGRES_DB -c "GRANT SELECT, UPDATE, INSERT,DELETE ON prices TO $POSTGRES_USER;"
go run . &