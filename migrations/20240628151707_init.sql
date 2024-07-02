-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS warehouses
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    is_available BOOLEAN      NOT NULL
);

CREATE TABLE IF NOT EXISTS products
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    code       INT          NOT NULL,
    quantity   INT          NOT NULL,
    size       INT          NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (code)
);

CREATE TABLE IF NOT EXISTS product_warehouse
(
    product_id        INT NOT NULL REFERENCES products (id),
    warehouse_id      INT NOT NULL REFERENCES warehouses (id),
    quantity          INT NOT NULL,
    reserved_quantity INT,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, warehouse_id)
);


-- CREATE TABLE IF NOT EXISTS reservations
-- (
--     product_id        INT NOT NULL REFERENCES products (id),
--     warehouse_id      INT NOT NULL REFERENCES warehouses (id),
--     reserved_quantity INT NOT NULL,
--     created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (product_id, warehouse_id)
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS reservations;
DROP TABLE IF EXISTS product_warehouse;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS warehouses;
-- +goose StatementEnd
