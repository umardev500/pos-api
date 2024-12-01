-- +goose Up
CREATE TABLE IF NOT EXISTS product_pricings (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    unit_id UUID NULL,
    custom_unit_id UUID NULL,
    tier_id UUID NOT NULL,
    -- Price tier name (e.g., "regular", "wholesale")
    price FLOAT NOT NULL,
    FOREIGN KEY (unit_id) REFERENCES units (id) ON DELETE CASCADE,
    FOREIGN KEY (custom_unit_id) REFERENCES custom_units (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (tier_id) REFERENCES tiers (id) ON DELETE CASCADE,
    CONSTRAINT chk_unit CHECK (
        (
            unit_id IS NOT NULL
            AND custom_unit_id IS NULL
        )
        OR (
            unit_id IS NULL
            AND custom_unit_id IS NOT NULL
        )
    )
);
-- Seed data
INSERT INTO product_pricings (
        id,
        product_id,
        unit_id,
        custom_unit_id,
        tier_id,
        price
    )
VALUES (
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000',
        NULL,
        '00000000-0000-0000-0000-000000000000',
        10.0
    ), (
        '00000000-0000-0000-0000-000000000001',
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000001',
        NULL,
        '00000000-0000-0000-0000-000000000001',
        5.0

    ), (
        '00000000-0000-0000-0000-000000000002',
        '00000000-0000-0000-0000-000000000000',
        NULL,
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000',
        15.0
    );
-- +goose Down
DROP TABLE product_pricings;