CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE "GENRES" (
  "id" uuid default uuid_generate_v4() primary key,
  "name" varchar not null,
  "created_at" timestamptz not null default (now())
);

CREATE TABLE "PUBLISHERS" (
  "id" uuid default uuid_generate_v4() primary key,
  "name" varchar not null,
  "address" varchar,
  "created_at" timestamptz not null default (now())
);

CREATE TABLE "AUTHORS" (
  "id" uuid default uuid_generate_v4() primary key, 
  "full_name" varchar not null,
  "birthdate" DATE,
  "created_at" timestamptz not null default (now())
);

CREATE TABLE "BOOKS" (
  "id" uuid default uuid_generate_v4() primary key,
  "title" varchar not null,
  "full_title" varchar not null,
  "publisher" uuid not null,
  "publication_date" date not null,
  "isbn" varchar not null,
  "description" varchar,
  "price" float not null,
  "stock_quantity" integer not null,
  "front_cover_image" varchar,
  "back_cover_image" varchar,
  "created_at" timestamptz not null default (now()),
  CONSTRAINT check_price CHECK(
    price > 0
  ),
  CONSTRAINT check_stock_quantity CHECK(
    stock_quantity > 0
  )
);

CREATE TABLE "BOOK_AUTHORS" (
  "book_id" uuid,
  "author_id" uuid,
  "created_at" timestamptz not null default (now()),
  primary key ("book_id", "author_id")
);

CREATE TABLE "BOOK_GENRES" (
  "book_id" uuid,
  "genre_id" uuid,
  "created_at" timestamptz not null default (now()),
  primary key ("book_id", "genre_id")
);

ALTER TABLE "BOOK_GENRES"
ADD CONSTRAINT fk_book_genres_books FOREIGN KEY ("book_id") REFERENCES "BOOKS" ("id");

ALTER TABLE "BOOK_GENRES"
ADD CONSTRAINT fk_book_genres_genres FOREIGN KEY ("genre_id") REFERENCES "GENRES" ("id");

ALTER TABLE "BOOK_AUTHORS"
ADD CONSTRAINT fk_book_authors_books FOREIGN KEY ("book_id") REFERENCES "BOOKS" ("id");

ALTER TABLE "BOOK_AUTHORS"
ADD CONSTRAINT fk_book_authors_authors FOREIGN KEY ("author_id") REFERENCES "AUTHORS" ("id");

ALTER TABLE "BOOKS"
ADD CONSTRAINT fk_books_publishers FOREIGN KEY ("publisher") REFERENCES "PUBLISHERS" ("id");
