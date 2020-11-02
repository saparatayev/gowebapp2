create table if not exists category(
    id serial primary key,
    description varchar(100) not NULL
);

create table if not EXISTS products(
	id BIGSERIAL PRIMARY KEY,
	NAME VARCHAR(255) NOT NULL,
	price REAL NOT NULL,
	quantity INTEGER DEFAULT 0,
	amount REAL DEFAULT 0.0,
	category BIGINT NOT NULL,
	CONSTRAINT products_category_fk FOREIGN KEY(category)
	REFERENCES category(id)
);
