CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "USERS" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar(11) NOT NULL,
  "full_name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'Active',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT check_user_status CHECK(
    "status" in ('Active', 'Blocked', 'Inactive')
  )
);