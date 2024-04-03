CREATE TABLE IF NOT EXISTS orders(
    id VARCHAR PRIMARY KEY NOT NULL,
    data JSON NOT NULL
);

-- INSERT INTO orders (id, data) VALUES (1, '{"first": 1}');
-- INSERT INTO orders (id, data) VALUES (2, '{"second": 2}');
