CREATE TABLE IF NOT EXISTS item (
  id          INTEGER PRIMARY KEY,
  filename    TEXT NOT NULL,
  contents    BLOB NOT NULL,
  filetype    TEXT,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
