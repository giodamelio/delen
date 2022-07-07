// Code generated by sqlc. DO NOT EDIT.
// source: items.sql

package db

import (
	"context"
	"database/sql"
)

const createItem = `-- name: CreateItem :one
INSERT INTO item (
  name, type, contents
) VALUES (
  $1, $2, $3
)
RETURNING id, name, type, contents
`

type CreateItemParams struct {
	Name     string
	Type     sql.NullString
	Contents []byte
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem, arg.Name, arg.Type, arg.Contents)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Contents,
	)
	return i, err
}

const listItems = `-- name: ListItems :many
SELECT id, name, type, contents FROM item
ORDER BY name
`

func (q *Queries) ListItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Contents,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}