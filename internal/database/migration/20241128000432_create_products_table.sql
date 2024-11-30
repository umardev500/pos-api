-- +goose Up
CREATE TABLE products (
    id UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    pricing JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- Create Index
CREATE INDEX idx_products_name ON products USING GIN("name" gin_trgm_ops);

-- Seed data with improved unit naming
INSERT INTO products (id, "name", "description", pricing)
VALUES
  (
    '00000000-0000-0000-0000-000000000000',  -- UUID for product
    'Eggs',  -- Product name
    'Fresh farm eggs',  -- Product description
    '[
        {
          "unit": {
            "name": "dozen",
            "multiplier": 12
          },
          "prices": [
            {
              "name": "regular",
              "price": 30.00
            },
            {
              "name": "wholesale",
              "price": 25.00
            }
          ]
        },
        {
          "unit": {
            "name": "pack",
            "multiplier": 10
          },
          "prices": [
            {
              "name": "regular",
              "price": 28.00
            },
            {
              "name": "wholesale",
              "price": 23.00
            }
          ]
        },
        {
          "unit": {
            "name": "kilogram",
            "multiplier": 1
          },
          "prices": [
            {
              "name": "regular",
              "price": 50.00
            },
            {
              "name": "wholesale",
              "price": 45.00
            }
          ]
        }
      ]
'::jsonb  -- JSONB format for pricing
  );
-- +goose Down
DROP TABLE products;