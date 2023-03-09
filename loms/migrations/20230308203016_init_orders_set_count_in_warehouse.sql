-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders_set_count_in_warehouse
(
    order_set_count_in_warehouse_id serial primary key,
    item_count                      int4   not null,
    warehouse_id                    bigint not null,
    order_set_id                    bigint not null
);

CREATE INDEX IF NOT EXISTS idx_warehouse_id ON orders_set_count_in_warehouse (warehouse_id);
CREATE INDEX IF NOT EXISTS idx_order_set_id ON orders_set_count_in_warehouse (order_set_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_warehouse_id;
DROP INDEX IF EXISTS idx_order_set_id;
DROP TABLE IF EXISTS orders_set_count_in_warehouse;
-- +goose StatementEnd
