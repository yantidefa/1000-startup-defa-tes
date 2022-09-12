-- +goose Up
-- +goose StatementBegin
CREATE TABLE "roles_points" (
    id          serial PRIMARY KEY        NOT NULL,
    r_point       int             NOT NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
);
-- +goose StatementEnd
DROP TYPE IF EXISTS "role";
CREATE TYPE "role" AS ENUM ('hacker','hustler','hipster');
ALTER TABLE "roles_points" ADD COLUMN "roles" "role";
-- +goose Down
-- +goose StatementBegin
DROP TABLE "roles_points";
-- +goose StatementEnd
