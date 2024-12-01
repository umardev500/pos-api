-- +goose Up
CREATE TABLE tiers (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- +goose Down
DROP TABLE tiers;
