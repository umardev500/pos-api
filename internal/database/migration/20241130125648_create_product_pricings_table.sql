-- +goose Up
CREATE TABLE IF NOT EXISTS product_pricings (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    unit_id UUID NULL,
    custom_unit_id UUID NULL,
    tier_id UUID NOT NULL, -- Price tier name (e.g., "regular", "wholesale")
    price FLOAT NOT NULL,

    FOREIGN KEY (unit_id) REFERENCES units (id) ON DELETE CASCADE,
    FOREIGN KEY (custom_unit_id) REFERENCES custom_units (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (tier_id) REFERENCES tiers (id) ON DELETE CASCADE,
    CONSTRAINT chk_unit CHECK (
        (unit_id IS NOT NULL AND custom_unit_id IS NULL) OR
        (unit_id IS NULL AND custom_unit_id IS NOT NULL)
    )
);

-- +goose Down
DROP TABLE product_pricings;
