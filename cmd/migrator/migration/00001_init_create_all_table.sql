-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text,
  "email" text,
  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "public"."users"(id, name, email) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858c8d', 'Hieu Phan', 'hieutrunglephan@gmail.com');
INSERT INTO "public"."users"(id, name, email) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858c8e', 'Hieu Dep Trai', 'hieudeptrai@gmail.com');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
