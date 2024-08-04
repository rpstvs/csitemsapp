-- name: GetPricebyID :one
SELECT Price
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC;
-- name: GetLatestPrice :one
SELECT Price,
    Item_id
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC;