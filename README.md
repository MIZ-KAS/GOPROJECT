## Download Migrate tool

 The migrate tool is used for database migrations.

https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g
https://github.com/golang-migrate/migrate

## Extra Work!! 
For you to create a migration, you need to create the kodboris database

Connect to the postgres database server through the docker container exec
and run the init.sql script in db/init.sql. You can confirm this by connecting to the 
kodboris database with the postgres user.

migrate create -ext sql -dir db/migration -seq  init-schema

## Copy db shcema into the migrate migration file

`Copy db schema from db/shcema.sql into *****_init-shecma.up.sql

## Run migrations

`migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" -verbose up`
