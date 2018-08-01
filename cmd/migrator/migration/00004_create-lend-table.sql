-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."lend_books" (
    "id" uuid PRIMARY KEY NOT NULL,
    "book_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "from" timestamptz DEFAULT now(),
    "to" timestamptz,
    "created_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz,
    CONSTRAINT "lend_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) NOT DEFERRABLE,
    CONSTRAINT "lend_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE
) WITH (oids = false);

INSERT INTO "public"."lend_books"(id, book_id, user_id) VALUES ('8ab72016-040d-40d9-9632-cf7a1e858a4e', '8ab72016-040d-40d9-9632-cf7a1e858c4e', '8ab72016-040d-40d9-9632-cf7a1e858c8d');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."lend_books";