-- +goose Up
-- +goose StatementBegin

CREATE TABLE "admins" (
    id          serial PRIMARY KEY        NOT NULL,
    name        character varying(30)     NOT NULL,
    birthday    date                      NULL,
    phone_no        character varying(30)     NOT NULL,
    id_domicile int         NOT NULL,
    image       character varying             NULL,
    email       character varying(30)     NOT NULL,
    password    character varying         NOT NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
    
);


ALTER TABLE "admins" ADD CONSTRAINT "admins_email" UNIQUE ("email");
ALTER TABLE "admins" ADD FOREIGN KEY ("id_domicile") REFERENCES "domiciles" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE admins DROP COLUMN IF EXISTS "gender";
DROP TABLE "admins";
-- +goose StatementEnd
