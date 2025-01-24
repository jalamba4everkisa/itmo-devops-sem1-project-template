#!/bin/bash
export PGPASSWORD=$POSTGRES_PASSWORD
sudo systemctl start postgresql
echo "POSTGRES_USER=$POSTGRES_USER" > .env
echo "POSTGRES_DB=$POSTGRES_DB" > .env
echo "POSTGRES_PASSWORD=$POSTGRES_PASSWORD" > .env
cat .env
psql -U $POSTGRES_USER -d $POSTGRES_DB -h $POSTGRES_HOST -c "CREATE TABLE prices(
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR NOT NULL,
                                        category VARCHAR NOT NULL,
                                        price REAL NOT NULL,
                                        create_date DATE NOT NULL
                                    );"
psql -U $POSTGRES_USER -d $POSTGRES_DB -h $POSTGRES_HOST -c "GRANT SELECT, UPDATE, INSERT,DELETE ON prices TO $POSTGRES_USER;"
psql -U $POSTGRES_USER -d $POSTGRES_DB -c -h $POSTGRES_HOST "GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public to $POSTGRES_USER;"
psql -U $POSTGRES_USER -d $POSTGRES_DB -h $POSTGRES_HOST -c "SELECT * FROM prices;"
./project_sem&