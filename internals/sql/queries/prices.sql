-- name: GetPricebyID :one
SELECT Price
FROM Prices
WHERE Item_id = $1;
-- name: GetBestItems :many
SELECT Item_id,
    Price,
    Items.ItemName,
    CAST (Items.DayChange AS NUMERIC(10, 2))
FROM Prices
    INNER JOIN Items ON Prices.Item_id = Items.Id
WHERE Items.DayChange IS NOT NULL
ORDER BY Items.DayChange DESC,
    PriceDate DESC
LIMIT 5;
-- name: GetWorstItems :many
SELECT Item_id,
    Price,
    Items.ItemName,
    CAST (Items.DayChange AS NUMERIC(10, 2))
FROM Prices
    INNER JOIN Items ON Prices.Item_id = Items.Id
ORDER BY Items.DayChange ASC,
    PriceDate DESC
LIMIT 5;