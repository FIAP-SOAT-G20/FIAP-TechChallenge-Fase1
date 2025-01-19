-- Inserindo categorias
INSERT INTO categories (name)
VALUES ('Lanches'),
       ('Bebidas'),
       ('Sobremesas'),
       ('Acompanhamentos'),
       ('Combos');

-- Inserindo funcionários
INSERT INTO staffs (name, role)
VALUES ('João Silva', 'COOK'),
       ('Maria Oliveira', 'ATTENDANT'),
       ('Pedro Santos', 'MANAGER'),
       ('Ana Costa', 'COOK'),
       ('Carlos Pereira', 'ATTENDANT');

-- Inserindo produtos
INSERT INTO products (name, description, price, category_id, image_url, staff_id, active)
VALUES ('X-Burger', 'Hambúrguer com queijo, alface e tomate', 25.90, 1, 'https://example.com/xburger.jpg', 1, true),
       ('Coca-Cola 350ml', 'Refrigerante Coca-Cola lata', 6.90, 2, 'https://example.com/coca.jpg', 2, true),
       ('Sundae', 'Sorvete com calda de chocolate', 12.90, 3, 'https://example.com/sundae.jpg', 1, true),
       ('Batata Frita', 'Porção de batata frita crocante', 15.90, 4, 'https://example.com/batata.jpg', 4, true),
       ('Combo Big', 'X-Burger + Batata + Refrigerante', 42.90, 5, 'https://example.com/combo.jpg', 1, true);

-- Inserindo clientes
INSERT INTO customers (name, email, cpf)
VALUES ('Lucas Mendes', 'lucas@email.com', '123.456.789-00'),
       ('Julia Santos', 'julia@email.com', '987.654.321-00'),
       ('Rafael Costa', 'rafael@email.com', '456.789.123-00'),
       ('Mariana Lima', 'mariana@email.com', '789.123.456-00'),
       ('Bruno Oliveira', 'bruno@email.com', '321.654.987-00');

-- Inserindo pedidos
INSERT INTO orders (customer_id, total_bill)
VALUES (1, 32.80),
       (2, 42.90),
       (3, 25.90),
       (4, 58.70),
       (5, 19.80);

-- Inserindo produtos dos pedidos
INSERT INTO order_products (order_id, product_id, price, quantity)
VALUES (1, 1, 25.90, 1),
       (1, 2, 6.90, 1),
       (2, 5, 42.90, 1),
       (3, 1, 25.90, 1),
       (4, 1, 25.90, 1),
       (4, 2, 6.90, 1),
       (4, 3, 12.90, 1),
       (4, 4, 15.90, 1),
       (5, 2, 6.90, 2),
       (5, 3, 12.90, 1);

-- Inserindo histórico dos pedidos
INSERT INTO order_histories (order_id, staff_id, status)
VALUES (1, 2, 'RECEIVED'),
       (1, 1, 'PREPARING'),
       (2, 2, 'READY'),
       (3, 2, 'COMPLETED'),
       (4, 2, 'RECEIVED'),
       (5, 2, 'PREPARING');

-- Inserindo pagamentos
INSERT INTO payments (id, status, external_payment_id, order_id, qr_data)
VALUES (1, 'CONFIRMED', '5c272292-4ba4-41e9-83d8-dea99afe5194', 1, 'QR_DATA_123'),
       (2, 'CONFIRMED', 'ac174c5e-c9ef-4407-a3b3-bceeb4163af3', 2, 'QR_DATA_456'),
       (3, 'PROCESSING', 'b7fa4bee-fc25-4bb4-b948-5139af948a39', 3, 'QR_DATA_789'),
       (4, 'CONFIRMED', '0bbd246e-c04d-44da-88f1-78b4b2cc354c', 4, 'QR_DATA_012'),
       (5, 'FAILED', '09d92b11-cd55-4a72-b2ee-7377ceefe265', 5, 'QR_DATA_345');

-- NEXT VAL
SELECT setval('categories_id_seq', (SELECT MAX(id) FROM categories));
SELECT setval('staffs_id_seq', (SELECT MAX(id) FROM staffs));
SELECT setval('products_id_seq', (SELECT MAX(id) FROM products));
SELECT setval('customers_id_seq', (SELECT MAX(id) FROM customers));
SELECT setval('orders_id_seq', (SELECT MAX(id) FROM orders));
SELECT setval('order_histories_id_seq', (SELECT MAX(id) FROM order_histories));