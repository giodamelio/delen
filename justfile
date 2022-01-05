set dotenv-load := false

@_list:
  just --list

# Start the server
go:
  cargo run

# Create a migration
sql-create-migration name:
  sqlx migrate add -r {{name}}

# Install devtools
install-tools:
  cargo install sqlx-cli