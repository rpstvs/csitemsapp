-- name: GetItemsIds :many
SELECT Id
FROM Items;
-- name: GetItemByName :one
SELECT Id
FROM Items
WHERE itemname = $1;
-- name: GetNameById :one
SELECT Itemname
FROM Items
WHERE Id = $1;
-- name: GetItemsRecord :many
SELECT Itemname,
    Prices.Price
FROM Items
    LEFT JOIN Prices ON Items.Id = Prices.Item_id
ORDER BY PriceDate DESC
Limit $1;
-- name: GetPriceHistory :one
SELECT CAST (Prices.Price AS NUMERIC(10, 2))
FROM Items
    LEFT JOIN Prices ON Items.Id = Prices.Item_id
WHERE Itemname = $1;
-- name: GetBestItems :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
WHERE DayChange IS NOT NULL
ORDER BY DayChange DESC
LIMIT 5;
-- name: GetWorstItems :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
ORDER BY DayChange ASC
LIMIT 5;
-- name: GetItemInfo :one
SELECT *
FROM Items
WHERE Itemname = $1;
-- name: GetItemSuggestion :many
SELECT ItemName
FROM Items
WHERE ItemName LIKE $1;