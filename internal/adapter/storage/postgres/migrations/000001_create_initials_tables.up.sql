CREATE TABLE IF NOT EXISTS customers
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    email      VARCHAR   NOT NULL UNIQUE,
    cpf        VARCHAR   NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS categories
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS staffs
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    role VARCHAR CHECK (role IN ('COOK', 'ATTENDANT', 'MANAGER'))
);

CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR        NOT NULL,
    description VARCHAR,
    price       DECIMAL(19, 2) NOT NULL,
    category_id INT            NOT NULL REFERENCES categories (id),
    image_url   VARCHAR,
    staff_id    INT REFERENCES staffs (id),
    active      BOOLEAN                 DEFAULT true,
    created_at  TIMESTAMP      NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP      NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS orders
(
    id          SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers (id),
    total_bill  DECIMAL(19, 2),
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS order_products
(
    order_id   INT REFERENCES orders (id),
    product_id INT REFERENCES products (id),
    price      DECIMAL(19, 2),
    quantity   INT NOT NULL,
    PRIMARY KEY (order_id, product_id)
);

CREATE TABLE IF NOT EXISTS payments
(
    id                  SERIAL PRIMARY KEY,
    status              VARCHAR CHECK (status IN ('PROCESSING', 'CONFIRMED', 'CANCELED', 'FAILED')) DEFAULT 'PROCESSING',
    external_payment_id VARCHAR,
    order_id            INT REFERENCES orders (id),
    qr_data             VARCHAR,
    created_at          TIMESTAMP NOT NULL                                             DEFAULT now(),
    updated_at          TIMESTAMP NOT NULL                                             DEFAULT now()
);