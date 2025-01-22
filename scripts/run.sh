#!/bin/bash
export PGPASSWORD=${POSTGRES_PASSWORD}
psql -U $POSTGRES_USER -d $POSTGRES_DB -h $POSTGRES_HOST -c "CREATE TABLE prices(
                                        id INTEGER PRIMARY KEY NOT NULL,
                                        name VARCHAR NOT NULL,
                                        category VARCHAR NOT NULL,
                                        price REAL NOT NULL,
                                        create_date DATE NOT NULL
                                    )";
go run .