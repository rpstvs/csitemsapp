-- +goose Up
CREATE TABLE Items (
    Id UUID PRIMARY KEY NOT NULL,
    ItemName TEXT UNIQUE NOT NULL
);
-- +goose Down
DROP TABLE Items;