BEGIN;

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    price NUMERIC(5,2)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    status VARCHAR(10) CHECK (status in ('not ready', 'ready', 'completed')) DEFAULT 'not ready'
);

CREATE TABLE items_orders (
    item_id INT REFERENCES items(id) NOT NULL,
    order_id INT REFERENCES orders(id) NOT NULL,
    CONSTRAINT id PRIMARY KEY (item_id, order_id),
    count INT NOT NULL
);

INSERT INTO items (name, price) VALUES ('Cheeseburger', 2.99), ('Hamburger', 1.50), ('Big Mac', 3.00);

COMMIT;
