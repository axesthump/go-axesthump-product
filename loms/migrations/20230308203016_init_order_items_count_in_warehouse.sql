-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items_count_in_warehouse
(
    id             serial primary key,
    count          int4   not null,
    warehouse_id   bigint not null,
    order_items_id bigint not null
);

CREATE INDEX IF NOT EXISTS idx_warehouse_id ON order_items_count_in_warehouse (warehouse_id);
CREATE INDEX IF NOT EXISTS idx_order_set_id ON order_items_count_in_warehouse (order_items_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_warehouse_id;
DROP INDEX IF EXISTS idx_order_set_id;
DROP TABLE IF EXISTS order_items_count_in_warehouse;
-- +goose StatementEnd
