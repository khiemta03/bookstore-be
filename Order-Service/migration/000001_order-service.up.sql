CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "DISCOUNTS" (
    discount_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    discount_code varchar(10) NOT NULL,
    discount_value float NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ORDERS" (
    order_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id varchar NOT NULL,
    order_at timestamptz NOT NULL DEFAULT (now()),
    status VARCHAR NOT NULL DEFAULT 'PENDING',
    discount UUID DEFAULT NULL,
    total_amount float NOT NULL,
    shipping_address VARCHAR NOT NULL,
    FOREIGN KEY (discount) REFERENCES Discounts(discount_id)
);

CREATE TABLE "ORDER_DETAILS" (
    order_id UUID NOT NULL,
    book_id varchar NOT NULL,
    quantity INT NOT NULL,
    unit_price float NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (order_id) REFERENCES Discounts(discount_id),
    PRIMARY KEY (order_id, book_id)
);

CREATE TABLE "SHOPPING_CART_ITEMS" (
    cart_item_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id varchar NOT NULL,
    book_id varchar NOT NULL,
    quantity INT NOT NULL,
    unit_price float NOT NULL,
    status VARCHAR NOT NULL DEFAULT 'ADDED',
    added_at timestamptz NOT NULL DEFAULT (now())
);