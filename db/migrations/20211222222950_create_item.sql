-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    type        TEXT,
    contents    BYTEA NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item;
-- +goose StatementEnd
