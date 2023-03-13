-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart_items
(
    id      serial primary key,
    sku     bigint not null,
    count   int4   not null,
    cart_id bigint not null
);

CREATE INDEX IF NOT EXISTS idx_id ON cart_items (cart_id);
CREATE INDEX IF NOT EXISTS idx_sku ON cart_items (sku);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_id;
DROP INDEX IF EXISTS idx_sku;
DROP TABLE IF EXISTS cart_items;
-- +goose StatementEnd