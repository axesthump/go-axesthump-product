-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS outbox_orders
(
    id       serial primary key,
    order_id bigint not null,
    status   int    not null,
    is_send  bool   not null default false
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS outbox_orders;
-- +goose StatementEnd