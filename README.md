# Expense Tracker REST API

## Steps
1. install postgreSQL on your local/docker (i'm using docker)
Docker : https://dev.to/shree_j/how-to-install-and-run-psql-using-docker-41j2

2. install golang

3. install project dependencies:
* Migrate will be use to create table
`https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#installation`
* SQLBoiler to execute script to database
`go get -u -t github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql`

4. To create table run this command
* assume that postgres is run & setup with `user=postgres` & `password=password`, and database `expense` exists.
`migrate -source file://internal/catalog/database/migrations -database "postgres://postgres:password@localhost:5432/expense?sslmode=disable" up`
tables will be created : expenses, users, category

5. to create models run this command
`(cd ./internal/catalog && sqlboiler --add-soft-deletes psql)`
