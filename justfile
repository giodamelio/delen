@_list:
  just --list

# Run the binary
go:
  go run .

# Create a new migration
migrate-create name:
  goose -dir ./migrations sqlite3 ./db.sqlite3 create {{name}} sql

# Run all the migrations
migrate-up:
  goose -dir ./migrations sqlite3 ./db.sqlite3 up

# Roll back all the migrations
migrate-down:
  goose -dir ./migrations sqlite3 ./db.sqlite3 down

# Reset the DB
db-reset:
  rm db.sqlite3
  @just migrate-up
