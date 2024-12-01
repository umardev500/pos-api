-- +goose Up
CREATE TABLE products (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- Create Index
CREATE INDEX idx_products_name ON products USING GIN("name" gin_trgm_ops);

-- Seed data with improved unit naming
INSERT INTO products (id, "name", "description")
VALUES
  (
    '00000000-0000-0000-0000-000000000000',  -- UUID for product
    'Eggs',  -- Product name
    'Fresh farm eggs'  -- Product description
  );
-- +goose Down
DROP TABLE products;