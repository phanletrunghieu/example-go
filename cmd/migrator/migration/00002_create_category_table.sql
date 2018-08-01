-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."categories" (
  "id" uuid PRIMARY KEY NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text
) WITH (oids = false);

INSERT INTO "public"."categories"(id, name) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858c9d', 'Bananas');
INSERT INTO "public"."categories"(id, name) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858c9e', 'Apple');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."categories";