-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders_set
(
    order_set_id serial primary key,
    item_sku     bigint not null,
    items_count  bigint not null,
    order_id     bigint not null
);

CREATE INDEX IF NOT EXISTS idx_items_count ON orders_set (items_count);
CREATE INDEX IF NOT EXISTS idx_order_id ON orders_set (order_id);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_items_count;
DROP INDEX IF EXISTS idx_order_id;
DROP TABLE IF EXISTS orders_set;
-- +goose StatementEnd
