-- +goose Up
-- +goose StatementBegin
CREATE TABLE "books" (
    id                 serial PRIMARY KEY NOT NULL,
    tittle             character varying(50)  NOT NULL,
    description        text  NOT NULL,
    book               character varying  NOT NULL,
    edition            character varying(100)  NOT NULL,
    publication_year   character varying(5)  NOT NULL,
    author             character varying (100)  NOT NULL,
    id_books_category  int  NOT NULL,
    created_at        timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         timestamp              NULL
);
-- +goose StatementEnd
ALTER TABLE "books" ADD FOREIGN KEY ("id_books_category") REFERENCES "books_categorys" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
-- +goose Down
-- +goose StatementBegin
DROP TABLE "books";
-- +goose StatementEnd
