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
-- Seed data
INSERT INTO custom_units (id, product_id, name, multiplier, created_by)
VALUES (
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000',
        'Crate Custom Unit',
        1,
        '00000000-0000-0000-0000-000000000000'
    );
-- +goose Down
DROP TABLE IF EXISTS custom_units;