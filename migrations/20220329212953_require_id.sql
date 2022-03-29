-- +goose Up
-- +goose StatementBegin

-- SQLite doesn't support changing column types, so we have to recreate and copy data
CREATE TABLE new_items (
  id INTEGER PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  mimeType TEXT NOT NULL,
  contents BLOB NOT NULL
);

-- Set the column mimeType to text/plain for any that are null
UPDATE items
SET mimeType = "text/plain"
WHERE mimeType IS null;

-- Copy data to new table
INSERT INTO new_items SELECT * FROM items;

-- Drop old table and rename the new table
DROP TABLE items;
ALTER TABLE new_items RENAME TO items;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- SQLite doesn't support changing column types, so we have to recreate and copy data
CREATE TABLE new_items (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  mimeType TEXT,
  contents BLOB
);

INSERT INTO new_items SELECT * FROM items;

DROP TABLE items;

ALTER TABLE new_items RENAME TO items;
-- +goose StatementEnd
