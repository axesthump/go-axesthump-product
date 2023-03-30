-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS outbox_orders
(
    id          serial primary key,
    order_id    bigint not null,
    status      int    not null,
    send_status int    not null default 1,
    err_message text   not null default ''
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS outbox_orders;
-- +goose StatementEnd