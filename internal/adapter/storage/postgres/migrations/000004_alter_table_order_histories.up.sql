CREATE TABLE IF NOT EXISTS order_histories
(
    id          SERIAL PRIMARY KEY ,
    order_id   INT REFERENCES orders (id) NOT NULL,
    staff_id   INT REFERENCES staffs (id) NULL,
    status     VARCHAR CHECK (status IN ('RECEIVED', 'PREPARING', 'READY', 'COMPLETED')) DEFAULT 'RECEIVED',
    created_at TIMESTAMP NOT NULL                                                        DEFAULT now(),
    updated_at TIMESTAMP NOT NULL                                                        DEFAULT now()
);