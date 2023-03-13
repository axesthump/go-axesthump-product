-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items
(
    id       serial primary key,
    sku      bigint not null,
    count    bigint not null,
    order_id bigint not null
);

CREATE INDEX IF NOT EXISTS idx_items_count ON order_items (count);
CREATE INDEX IF NOT EXISTS idx_order_id ON order_items (order_id);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_items_count;
DROP INDEX IF EXISTS idx_order_id;
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
