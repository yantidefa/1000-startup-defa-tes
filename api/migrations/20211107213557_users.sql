-- +goose Up
-- +goose StatementBegin

CREATE TABLE "users" (
    id          serial PRIMARY KEY        NOT NULL,
    name        character varying(30)     NOT NULL,
    birthday    date                      NULL,
    phone_no        character varying(30)     NOT NULL,
    id_startup int        NOT NULL,
    position     character varying         NOT NULL,
    id_domicile int         NOT NULL,
    image       character varying             NULL,
    email       character varying(30)     NOT NULL,
    password    character varying         NOT NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
    
);

DROP TYPE IF EXISTS "genderlist";
CREATE TYPE "genderlist" AS ENUM ('male','female');
ALTER TABLE "users" ADD COLUMN "gender" "genderlist";
ALTER TABLE "users" ADD CONSTRAINT "users_email" UNIQUE ("email");
ALTER TABLE "users" ADD FOREIGN KEY ("id_domicile") REFERENCES "domiciles" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE "users" ADD FOREIGN KEY ("id_startup") REFERENCES "startups" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE users DROP COLUMN IF EXISTS "gender";
DROP TABLE "users";
-- +goose StatementEnd
