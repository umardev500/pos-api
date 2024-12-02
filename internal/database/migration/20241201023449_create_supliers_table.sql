-- +goose Up
CREATE TABLE IF NOT EXISTS suppliers (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    "address" VARCHAR(255) NULL,
    "email" VARCHAR(50) NULL,
    "phone" VARCHAR(20) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);

-- +goose Down
DROP TABLE IF EXISTS suppliers;
