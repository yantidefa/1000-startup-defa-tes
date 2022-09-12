-- +goose Up
-- +goose StatementBegin

CREATE TABLE "blogs" (
    id                  serial PRIMARY KEY        NOT NULL,
    tittle              character varying(50)     NOT NULL,
    id_blog_category    int   NOT NULL,
    content             text        NOT NULL,
    image               character varying             NULL,
    created_at          timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          timestamp NULL
    
);

ALTER TABLE "blogs" ADD FOREIGN KEY ("id_blog_category") REFERENCES "blogs_categories" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE blogs DROP COLUMN IF EXISTS "gender";
DROP TABLE "blogs";
-- +goose StatementEnd
