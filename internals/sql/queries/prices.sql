-- name: GetPricebyID :one
SELECT Price
FROM Prices
WHERE Item_id = $1;
-- name: GetBestItems :many
SELECT Item_id,
    Price
FROM Prices
ORDER BY Price DESC
LIMIT 5;