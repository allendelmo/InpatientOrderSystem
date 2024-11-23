-- +goose Up
CREATE TABLE statuses (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE medication_orders (
    order_number SERIAL PRIMARY KEY,
    file_number INTEGER NOT NULL,
    nurse_name TEXT,
    ward TEXT,
    bed TEXT,
    medication TEXT,
    quantity INTEGER,
    uom TEXT,
    request_time TIMESTAMP NOT NULL,
    nurse_remarks TEXT,
    status_id UUID NOT NULL REFERENCES statuses(id),
    pharmacy_remarks TEXT
);
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    ward TEXT NOT NULL,
    permission TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    first_name TEXT,
    last_name TEXT
);
-- +goose Down
DROP TABLE medication_orders,
users,
statuses;