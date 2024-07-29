-- name: GetPricebyID :one
SELECT Price
FROM Prices
WHERE Item_id = $1;
-- name: GetBestItems :many
SELECT Item_id,
    Price,
    Items.ItemName
FROM Prices
    LEFT JOIN Items ON Prices.Item_id = Items.Id
ORDER BY Price DESC,
    PriceDate DESC
LIMIT 5;
-- name: GetWorstItems :many
SELECT Item_id,
    Price,
    Items.ItemName
FROM Prices
    LEFT JOIN Items ON Prices.Item_id = Items.Id
ORDER BY Price ASC,
    PriceDate DESC
LIMIT 5;