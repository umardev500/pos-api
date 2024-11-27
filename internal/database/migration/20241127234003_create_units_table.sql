-- +goose Up
CREATE TABLE units (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "description" VARCHAR(255),
    multiplier FLOAT NOT NULL
);
-- Seed data
INSERT
INSERT INTO units (id, "name", multiplier)
VALUES (uuid_generate_v4(), 'Piece', 1),
    (uuid_generate_v4(), 'Kilogram', 1000),
    (uuid_generate_v4(), 'Gram', 1);
-- +goose Down
DROP TABLE units;