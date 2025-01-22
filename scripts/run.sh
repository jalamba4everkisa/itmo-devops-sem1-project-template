#!/bin/bash
export PGPASSWORD=$POSTGRES_PASSWORD
sudo bash -c "echo -e 'local   all             all                                  trust' > /etc/postgresql/15/main/pg_hba.conf"
sudo bash -c "echo -e 'host   $POSTGRES_DB           $POSTGRES_USER          0.0.0.0/0            md5' >> /etc/postgresql/15/main/pg_hba.conf"
sudo systemctl start postgresql
echo $POSTGRES_DB
sudo -u postgres psql -c 'CREATE DATABASE "'$POSTGRES_DB'";'
#psql -U postgres -d $POSTGRES_DB -c "CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';"
#psql -U postgres -d $POSTGRES_DB -c "CREATE TABLE prices(
#                                        id INTEGER PRIMARY KEY NOT NULL,
#                                        name VARCHAR NOT NULL,
#                                        category VARCHAR NOT NULL,
#                                        price REAL NOT NULL,
#                                        create_date DATE NOT NULL
#                                    )";
#psql -U postgres -d $POSTGRES_DB -c "GRANT SELECT, UPDATE, INSERT,DELETE ON prices TO $POSTGRES_USER;"
go run . &