@_list:
    just --list

# Run the app
run:
    go run ./...

# Run all the migrations
migrate:
    goose -dir db/migrations/ postgres "user=postgres dbname=app sslmode=disable" up

# Create a new migration
create-migration name:
    goose -dir db/migrations/ postgres "user=postgres dbname=app sslmode=disable" create {{name}} sql

# Open a SQL console
sql-console:
    usql pg://postgres@localhost/app?sslmode=disable

# Generate the SQL Go code with SQLC
sql-generate:
    sqlc generate