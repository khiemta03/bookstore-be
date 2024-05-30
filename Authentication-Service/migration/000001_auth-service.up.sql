CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "SESSIONS" (
    "session_id" uuid PRIMARY KEY,
    "user_id" varchar NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "status" varchar NOT NULL DEFAULT 'Active',
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT check_session_status CHECK(
    "status" in ('Active', 'Blocked')
  )
);

CREATE TABLE "ACCESS_TOKEN" (
    "access_token_id" uuid PRIMARY KEY,
    "session_id" uuid NOT NULL,
    "access_token_value" varchar NOT NULL,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);