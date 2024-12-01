-- +goose Up
CREATE TABLE IF NOT EXISTS product_stocks (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    quantity FLOAT NOT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS product_stocks;
