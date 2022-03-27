DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT,
  email TEXT
);

INSERT INTO users (name,email)
VALUES
  ("Hayfa King","non.bibendum@icloud.ca"),
  ("Pascale Roberson","dignissim.magna@outlook.couk"),
  ("Kane Jacobs","molestie.sed@outlook.edu"),
  ("Lana Tillman","in.dolor.fusce@aol.ca"),
  ("Nasim Love","duis@aol.edu");
