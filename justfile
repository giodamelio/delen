@_list:
    just --list

migrate:
    migrate -source file://db/migrations/ -database sqlite://db/dev.sqlite up

create-migration name:
    migrate create -dir ./db/migrations/ -seq -ext sql {{name}}

install-tools:
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest