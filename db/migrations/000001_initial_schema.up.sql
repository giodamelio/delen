CREATE TABLE IF NOT EXISTS item (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    type        TEXT,
    contents    BLOB NOT NULL
);