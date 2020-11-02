INSERT INTO category(description) VALUES ('Fruits');
INSERT INTO category(description) VALUES ('Books');
INSERT INTO category(description) VALUES ('Cars');
INSERT INTO category(description) VALUES ('Informatics');

INSERT INTO products(name, price, quantity, amount, category) VALUES ('Laranja', 1.80, 100, (1.8 * 100), 1);
INSERT INTO products(name, price, quantity, amount, category) VALUES ('Banana', 1.10, 80, (1.8 * 80), 1);

INSERT INTO products(name, price, quantity, amount, category) VALUES ('Harry Potter', 20, 10, (10 * 10), 2);
INSERT INTO products(name, price, quantity, amount, category) VALUES ('Game of Thrones', 40, 20, (40 * 20), 2);

INSERT INTO products(name, price, quantity, amount, category) VALUES ('Ferrari', 900000, 5, (900000 * 5), 3);
INSERT INTO products(name, price, quantity, amount, category) VALUES ('Jaguar F-Type', 600000, 2, (600000 * 2), 3);

INSERT INTO products(name, price, quantity, amount, category) VALUES ('Computer', 3000, 12, (3000 * 12), 4);
INSERT INTO products(name, price, quantity, amount, category) VALUES ('Notebook', 2500, 8, (2500 * 8), 4);


SELECT * FROM products INNER JOIN category ON category.id = products.category;
SELECT category.description, products.* FROM products INNER JOIN category ON category.id = products.category;
