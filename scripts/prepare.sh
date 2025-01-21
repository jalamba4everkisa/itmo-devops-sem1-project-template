#!/bin/bash
sudo apt update -y
sudo apt install wget -y

wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
rm go1.23.5.linux-amd64.tar.gz
go get github.com/jackc/pgx/v5
go get github.com/joho/godotenv
sudo apt install postgresql -y
export $(grep -v '^#' .env | xargs) 
sudo -u postgres psql -c 'CREATE DATABASE "'$DB_PG'";'
sudo -u postgres psql -d $DB_PG -c "CREATE USER $USER_PG WITH PASSWORD '$PASSWORD_PG';"
sudo -u postgres psql -d $DB_PG -c "GRANT ALL PRIVILEGES ON DATABASE $DB_PG to $USER_PG;"
sudo -u postgres psql -d $DB_PG -c "CREATE TABLE prices(
                                        id INTEGER PRIMARY KEY NOT NULL,
                                        name VARCHAR NOT NULL,
                                        category VARCHAR NOT NULL,
                                        price REAL NOT NULL,
                                        create_date DATE NOT NULL
                                    )";
sudo -u postgres psql -d $DB_PG -c "GRANT SELECT, UPDATE, INSERT,DELETE ON prices TO $USER_PG;"
sudo bash -c "echo -e 'host    all             all             0.0.0.0/0            md5' > /etc/postgresql/16/main/pg_hba.conf"
sudo systemctl restart postgresql