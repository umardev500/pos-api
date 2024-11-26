-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(250) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Create Index
CREATE INDEX idx_users_username ON users USING GIN(username gin_trgm_ops);
-- Seed data
INSERT INTO users (id, username, email, password_hash)
VALUES (
        '00000000-0000-0000-0000-000000000000',
        'username1',
        'username1@example.com',
        '$2a$10$fRypLiLtPOrgB2Mvy6Fx5.oZuNwock3V3cVMTAj1wyn.Paam2Oeyu '
    ),
    (
        '00000000-0000-0000-0000-000000000001',
        'username2',
        'username2@example.com',
        '$2a$10$fRypLiLtPOrgB2Mvy6Fx5.oZuNwock3V3cVMTAj1wyn.Paam2Oeyu '
    );
-- +goose Down
DROP TABLE users;