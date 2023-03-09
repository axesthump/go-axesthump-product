-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS warehouses
(
    warehouse_id serial primary key
);

INSERT INTO warehouses(warehouse_id)
values (1),
       (2),
       (3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS warehouses;
-- +goose StatementEnd
