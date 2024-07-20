-- name: GetPricebyID :one
SELECT Price
FROM Prices
WHERE Item_id = $1;