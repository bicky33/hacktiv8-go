CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY NOT NULL,
    customer_name VARCHAR(255), 
    ordered_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS items (
    item_id SERIAL PRIMARY KEY NOT NULL,
    item_code VARCHAR(255) NOT NULL, 
    description VARCHAR(255) NOT NULL, 
    quantity INTEGER DEFAULT 0, 
    order_id INTEGER,
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
);