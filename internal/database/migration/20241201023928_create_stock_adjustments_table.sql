-- +goose Up
CREATE TABLE IF NOT EXISTS stock_adjustments (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    supplier_id UUID NULL, -- Nullable for non-supplier adjustments
    quantity INT NOT NULL,
    reason VARCHAR(255) NOT NULL, -- Reason for the adjustment
    reversed BOOLEAN NOT NULL DEFAULT FALSE,
    reversal_of UUID NULL, -- Link to the reversal adjustment
    created_by UUID NOT NULL, -- User who created the adjustment
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (supplier_id) REFERENCES suppliers (id) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE CASCADE,
    -- Reversal
    FOREIGN KEY (reversal_of) REFERENCES stock_adjustments (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS stock_adjustments;
