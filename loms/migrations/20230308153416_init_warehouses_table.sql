-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS warehouses
(
    id serial primary key
);

INSERT INTO warehouses(id)
values (1),
       (2),
       (3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS warehouses;
-- +goose StatementEnd
