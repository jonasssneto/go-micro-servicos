CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	product_name VARCHAR(50) NOT NULL,
	price NUMERIC(10,2) NOT NULL
);

CREATE TABLE orders (
	id SERIAL PRIMARY KEY,
	product_id INT NOT NULL,
	order_date DATE NOT NULL,
	order_status character varying NOT NULL DEFAULT 'pending',
	FOREIGN KEY (product_id) REFERENCES product(id)
);