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
SELECT CAST (Prices.Price AS NUMERIC(10, 2)),
    Prices.PriceDate
FROM Items
    LEFT JOIN Prices ON Items.Id = Prices.Item_id
WHERE Itemname = $1
LIMIT 7;
-- name: GetBestItems :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    CAST (WeekChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
WHERE DayChange < 40
ORDER BY DayChange DESC
LIMIT 5;
-- name: GetWorstItems :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
WHERE DayChange > -40
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
-- name: GetBestItemsWeek :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    CAST (WeekChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
WHERE WeekChange < 20
ORDER BY WeekChange DESC
LIMIT 5;
-- name: GetWorstItemsWeek :many
SELECT ItemName,
    Id,
    CAST (DayChange AS NUMERIC(10, 2)),
    CAST (WeekChange AS NUMERIC(10, 2)),
    ImageUrl
FROM Items
WHERE WeekChange > -20
ORDER BY DayChange DESC
LIMIT 5;