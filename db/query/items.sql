-- name: ListItems :many
SELECT * FROM item
ORDER BY name;

-- name: CreateItem :one
INSERT INTO item (
  name, type, contents
) VALUES (
  $1, $2, $3
)
RETURNING *;