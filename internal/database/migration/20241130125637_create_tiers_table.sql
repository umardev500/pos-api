-- +goose Up
CREATE TABLE tiers (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);

-- Seed data
INSERT INTO tiers (id, "name")
VALUES ('00000000-0000-0000-0000-000000000000', 'Regular'),
    ('00000000-0000-0000-0000-000000000001', 'Wholesale');

-- +goose Down
DROP TABLE tiers;
