CREATE TYPE order_status AS ENUM ('OPEN','CANCELLED','PENDING','RECEIVED', 'PREPARING', 'READY', 'COMPLETED');

alter table orders
    ALTER status DROP DEFAULT,
    ALTER COLUMN status TYPE order_status USING status::order_status,
    ALTER COLUMN status SET DEFAULT 'OPEN';

alter table order_histories
    ALTER status DROP DEFAULT,
    ALTER COLUMN status TYPE order_status USING status::order_status,
    ALTER COLUMN status SET DEFAULT 'OPEN';

ALTER TABLE orders DROP CONSTRAINT orders_status_check;
ALTER TABLE order_histories DROP CONSTRAINT order_histories_status_check;
