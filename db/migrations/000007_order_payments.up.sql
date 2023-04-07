
CREATE TYPE stat AS ENUM ('paid', 'waiting', 'unpaid');
CREATE TABLE IF NOT EXISTS order_payments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    amount bigint,
    status stat,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    order_id bigint NOT NULL    
);