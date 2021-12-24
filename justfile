@_list:
    just --list

migrate:
    goose -dir db/migrations/ postgres "user=postgres dbname=app sslmode=disable" up

create-migration name:
    goose -dir db/migrations/ postgres "user=postgres dbname=app sslmode=disable" create {{name}} sql

sql:
    usql pg://postgres@localhost/app?sslmode=disable

install-tools:
    go install github.com/pressly/goose/v3/cmd/goose@latest
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    go install github.com/xo/usql@master