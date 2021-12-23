-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    type        TEXT,
    contents    BLOB NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item;
-- +goose StatementEnd
