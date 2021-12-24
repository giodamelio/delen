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

# Install the CLI tools needed for the project
install-tools:
    go install github.com/pressly/goose/v3/cmd/goose@latest
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    go install github.com/xo/usql@master