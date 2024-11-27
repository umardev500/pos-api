-- +goose Up
CREATE TABLE units (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    multiplier FLOAT NOT NULL
);
-- Seed data
INSERT INTO units (id, "name", multiplier)
VALUES ("00000000-0000-0000-0000-000000000000", 'Piece', 1),
    ("00000000-0000-0000-0000-000000000001", 'Kilogram', 1000),
    ("00000000-0000-0000-0000-000000000002", 'Gram', 1);
-- +goose Down
DROP TABLE units;