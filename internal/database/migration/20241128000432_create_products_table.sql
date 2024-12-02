-- +goose Up
CREATE TABLE products (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    category_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);

-- Create Index
CREATE INDEX idx_products_name ON products USING GIN("name" gin_trgm_ops);
CREATE INDEX idx_product_deleted_at ON products (deleted_at);

-- Seed data with improved unit naming
INSERT INTO products (id, "name", "description", category_id)
VALUES
  (
    '00000000-0000-0000-0000-000000000000',  -- UUID for product
    'Eggs',  -- Product name
    'Fresh farm eggs',  -- Product description
    '00000000-0000-0000-0000-000000000000'  -- Category ID
  );
-- +goose Down
DROP TABLE products;