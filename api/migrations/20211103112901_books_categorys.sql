-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE "books_categories" (
    id          serial PRIMARY KEY        NOT NULL,
    name        character varying(30)      NOT   NULL,
    description        text      NOT   NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "books_categories";
-- +goose StatementEnd
