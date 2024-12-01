-- +goose Up
CREATE TABLE IF NOT EXISTS product_stocks (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- Seed data
INSERT INTO product_stocks (id, product_id, quantity)
VALUES ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 10);

-- +goose Down
DROP TABLE IF EXISTS product_stocks;
