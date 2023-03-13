-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
    id serial primary key,
    user_id bigint not null
);

CREATE INDEX IF NOT EXISTS idx_user_id ON carts(user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_user_id;
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd
