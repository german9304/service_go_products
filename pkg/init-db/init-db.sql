
--- Create products table if not exists
CREATE TABLE IF NOT EXISTS products (
	id varchar(30),
	name varchar(30),
	price integer
);


INSERT INTO products (id, name, price) VALUES ('1234', 'bike', 200);
