@_list:
    just --list

migrate:
    goose -dir db/migrations/ sqlite ./db/dev.sqlite up

create-migration name:
    goose -dir db/migrations/ sqlite create {{name}} sql

install-tools:
    go install github.com/pressly/goose/v3/cmd/goose@latest
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    go install github.com/xo/usql@master