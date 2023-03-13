-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
    id serial primary key,
    status   int    not null,
    user_id  bigint not null
);

CREATE INDEX IF NOT EXISTS idx_user_id ON orders (user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_user_id;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
