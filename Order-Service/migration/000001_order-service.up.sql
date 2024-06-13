CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "DISCOUNTS" (
    discount_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    discount_code varchar(10) NOT NULL,
    discount_value float NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT check_discount_duration CHECK(
        start_date < end_date
    )
);

CREATE TABLE "ORDERS" (
    order_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id varchar NOT NULL,
    order_at timestamptz NOT NULL DEFAULT (now()),
    status VARCHAR NOT NULL DEFAULT 'PENDING',
    discount UUID DEFAULT NULL,
    shipping_address VARCHAR NOT NULL,
    FOREIGN KEY (discount) REFERENCES "DISCOUNTS"(discount_id),
    CONSTRAINT check_status CHECK(
        status IN ('PENDING', 'PAYED', 'SHIPPED', 'PROCESSING')
    )
);

CREATE TABLE "ORDER_DETAILS" (
    order_id UUID NOT NULL,
    book_id varchar NOT NULL,
    quantity INT NOT NULL,
    unit_price float NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (order_id) REFERENCES "ORDERS"(order_id),
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
    CONSTRAINT check_status CHECK(
        status IN ('ADDED', 'REMOVED', 'ORDERED')
    )
);
CREATE UNIQUE INDEX unique_user_book_added ON "SHOPPING_CART_ITEMS" (user_id, book_id) WHERE status = 'ADDED';