-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts_items
(
    cart_item_id serial primary key,
    item_sku bigint not null,
    count int4 not null,
    cart_id bigint not null
);

CREATE INDEX IF NOT EXISTS idx_cart_id ON carts_items(cart_id);
CREATE INDEX IF NOT EXISTS idx_item_sku ON carts_items(item_sku);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_cart_id;
DROP INDEX IF EXISTS idx_item_sku;
DROP TABLE IF EXISTS carts_items;
-- +goose StatementEnd