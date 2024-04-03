CREATE TABLE IF NOT EXISTS racks
(
    id        SERIAL PRIMARY KEY,
    number    varchar
);

CREATE TABLE IF NOT EXISTS products
(
    id    SERIAL PRIMARY KEY,
    title varchar
);

CREATE TABLE racks_products
(
    id         SERIAL PRIMARY KEY,
    rack_id    int REFERENCES racks (id),
    product_id int REFERENCES products (id),
    priority varchar DEFAULT 'primary'
);

CREATE TABLE IF NOT EXISTS orders
(
    id         SERIAL PRIMARY KEY,
    created_at timestamp DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS order_details
(
    id         SERIAL PRIMARY KEY,
    order_id   int REFERENCES orders (id),
    product_id int REFERENCES products (id),
    count      int
);