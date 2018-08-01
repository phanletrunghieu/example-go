-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."books" (
  "id" uuid PRIMARY KEY NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text,
  "category_id" uuid NOT NULL,
  "author" text,
  "description" text,
  CONSTRAINT "books_category_id_fkey" FOREIGN KEY (category_id) REFERENCES categories(id) NOT DEFERRABLE
) WITH (oids = false);

INSERT INTO "public"."books"(id, name, category_id) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858c4e', 'Apple is not apple', '8ab72016-040d-40d9-9632-cf7a1e858c9e');
INSERT INTO "public"."books"(id, name, category_id) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858c4f', 'Banana is chuoi', '8ab72016-040d-40d9-9632-cf7a1e858c9d');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."books";