-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS warehouses_items
(
    sku          bigint,
    available    int4 not null,
    reserved     int4 not null,
    bought       int4 not null,
    warehouse_id bigint,
    PRIMARY KEY (sku, warehouse_id)
);

CREATE INDEX IF NOT EXISTS idx_item_sku ON warehouses_items (sku);
CREATE INDEX IF NOT EXISTS idx_warehouse_id ON warehouses_items (warehouse_id);

INSERT INTO warehouses_items (sku, available, reserved, bought, warehouse_id)
VALUES (1076963, 3, 0, 0, 1),
       (1148162, 3, 0, 0, 1),
       (1625903, 3, 0, 0, 1),
       (1076963, 3, 0, 0, 2),
       (1148162, 3, 0, 0, 2),
       (1625903, 3, 0, 0, 2),
       (1076963, 3, 0, 0, 3),
       (1148162, 3, 0, 0, 3),
       (1625903, 3, 0, 0, 3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_item_sku;
DROP INDEX IF EXISTS idx_warehouse_id;
DROP TABLE IF EXISTS warehouses_items;
-- +goose StatementEnd
