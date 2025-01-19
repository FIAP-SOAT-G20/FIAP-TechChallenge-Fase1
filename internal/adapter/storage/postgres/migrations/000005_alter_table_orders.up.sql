CREATE TYPE order_status AS ENUM ('OPEN','CANCELLED','PENDING','RECEIVED', 'PREPARING', 'READY', 'COMPLETED');

alter table orders
    ALTER status DROP DEFAULT,
    ALTER COLUMN status TYPE order_status USING status::order_status,
    ALTER COLUMN status SET DEFAULT 'OPEN';

alter table order_histories
    ALTER status DROP DEFAULT,
    ALTER COLUMN status TYPE order_status USING status::order_status,
    ALTER COLUMN status SET DEFAULT 'OPEN';

