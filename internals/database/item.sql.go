// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: item.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getItemByName = `-- name: GetItemByName :one
SELECT Id
FROM Items
WHERE itemname = $1
`

func (q *Queries) GetItemByName(ctx context.Context, itemname string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getItemByName, itemname)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getItemsIds = `-- name: GetItemsIds :many
SELECT Id
FROM Items
`

func (q *Queries) GetItemsIds(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getItemsIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNameById = `-- name: GetNameById :one
SELECT Itemname
FROM Items
WHERE Id = $1
`

func (q *Queries) GetNameById(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getNameById, id)
	var itemname string
	err := row.Scan(&itemname)
	return itemname, err
}
