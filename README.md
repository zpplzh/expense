# Expense Tracker REST API

## Steps
1. install postgreSQL on your local/docker (i'm using docker)\
Docker : https://dev.to/shree_j/how-to-install-and-run-psql-using-docker-41j2

2. install golang

3. install project dependencies:
* Migrate

`https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#installation`

* SQLBoiler

`go get -u -t github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql`

4. To create table run this command
* assume that postgres is run & setup with `user=postgres` & `password=password`, and database `expense` exists.

`migrate -source file://internal/app/database/migrations -database "postgres://postgres:password@localhost:5432/expense?sslmode=disable" up`

tables will be created : expenses, users, category

5. to create models run this command
`(cd ./internal/app && sqlboiler --add-soft-deletes psql)`

6. to run application export following environment value\
`export PORT = :[port]`  
`export DBNAME = "[database name]"`  
`export DBUSER = "[database user]"`  
`export DBPASS = "[databse pass]"`  

7. then run this command 
`go run ./cmd/expense`

# Usage

Expense Tracker REST API with several functionality such as:

## User
| API  | Description |
| ------------- | ------------- |
| POST /signup  | to register  |
| POST /login  | to login  |
| POST /logout  | to logout  |



## Category
| API  | Description |
| ------------- | ------------- |
| POST /categories  | to add category with icon  |
| GET /categories/{id} | to view category created before  |
| DELETE /categories/{id} | to delete category created before  |
| GET /categories | to view all category created before  |
| POST /categories/{id} | to update category created before  |

## Expense
| API  | Description |
| ------------- | ------------- |
| POST /expenses | to add expense |
| GET /expenses/{id} | to view expense created before  |
| GET /expenses | to view all expenses created before  |
| DELETE /expenses/{id} | to delete expenses created before  |
| UPDATE /expenses/{id} | to update expenses created before  |
| POST /batch/expenses | to add expense in batch |





