-- +goose Up
CREATE TABLE IF NOT EXISTS custom_units (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    multiplier FLOAT NOT NULL CHECK (multiplier > 0),
    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS custom_units;
