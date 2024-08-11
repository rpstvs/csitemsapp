// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: item.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getBestItems = `-- name: GetBestItems :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
WHERE DayChange IS NOT NULL
ORDER BY DayChange DESC
LIMIT 5
`

type GetBestItemsRow struct {
	Itemname  string
	ID        uuid.UUID
	Daychange float64
	Imageurl  string
}

func (q *Queries) GetBestItems(ctx context.Context) ([]GetBestItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getBestItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBestItemsRow
	for rows.Next() {
		var i GetBestItemsRow
		if err := rows.Scan(
			&i.Itemname,
			&i.ID,
			&i.Daychange,
			&i.Imageurl,
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

const getItemInfo = `-- name: GetItemInfo :one
SELECT id, itemname, daychange, weekchange, imageurl
FROM Items
WHERE Itemname = $1
`

func (q *Queries) GetItemInfo(ctx context.Context, itemname string) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemInfo, itemname)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Itemname,
		&i.Daychange,
		&i.Weekchange,
		&i.Imageurl,
	)
	return i, err
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

const getItemsRecord = `-- name: GetItemsRecord :many
SELECT Itemname,
    Prices.Price
FROM Items
    LEFT JOIN Prices ON Items.Id = Prices.Item_id
ORDER BY PriceDate DESC
Limit $1
`

type GetItemsRecordRow struct {
	Itemname string
	Price    sql.NullString
}

func (q *Queries) GetItemsRecord(ctx context.Context, limit int32) ([]GetItemsRecordRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemsRecord, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsRecordRow
	for rows.Next() {
		var i GetItemsRecordRow
		if err := rows.Scan(&i.Itemname, &i.Price); err != nil {
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

const getPriceHistory = `-- name: GetPriceHistory :one
SELECT CAST (Prices.Price AS NUMERIC(10, 2))
FROM Items
    LEFT JOIN Prices ON Items.Id = Prices.Item_id
WHERE Itemname = $1
`

func (q *Queries) GetPriceHistory(ctx context.Context, itemname string) (float64, error) {
	row := q.db.QueryRowContext(ctx, getPriceHistory, itemname)
	var prices_price float64
	err := row.Scan(&prices_price)
	return prices_price, err
}

const getWorstItems = `-- name: GetWorstItems :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
ORDER BY DayChange ASC
LIMIT 5
`

type GetWorstItemsRow struct {
	Itemname  string
	ID        uuid.UUID
	Daychange float64
	Imageurl  string
}

func (q *Queries) GetWorstItems(ctx context.Context) ([]GetWorstItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getWorstItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWorstItemsRow
	for rows.Next() {
		var i GetWorstItemsRow
		if err := rows.Scan(
			&i.Itemname,
			&i.ID,
			&i.Daychange,
			&i.Imageurl,
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
