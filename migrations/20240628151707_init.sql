-- +goose Up
-- +goose StatementBegin
CREATE TABLE warehouses
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    is_available BOOLEAN      NOT NULL
);

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    size        VARCHAR(50)  NOT NULL,
    unique_code VARCHAR(255) NOT NULL UNIQUE,
    quantity    INT
);

CREATE TABLE reservations
(
    id                SERIAL PRIMARY KEY,
    product_id        INT NOT NULL REFERENCES products (id),
    warehouse_id      INT NOT NULL REFERENCES warehouses (id),
    reserved_quantity INT NOT NULL CHECK (reserved_quantity >= 0),
    UNIQUE (product_id, warehouse_id)
);

CREATE TABLE reservation_warehouse
(
    reservation_id INT NOT NULL REFERENCES reservations (id),
    warehouse_id   INT NOT NULL REFERENCES warehouses (id),
    PRIMARY KEY (reservation_id, warehouse_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reservation_warehouse;
DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS warehouses;
-- +goose StatementEnd
