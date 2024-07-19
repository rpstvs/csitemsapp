-- name: GetItemsIds :many
SELECT Id
FROM Items;
-- name: GetItemByName :one
SELECT Id
FROM Items
WHERE itemname = $1;