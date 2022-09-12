-- +goose Up
-- +goose StatementBegin

CREATE TABLE "domiciles" (
    id          serial PRIMARY KEY        NOT NULL,
    city        character varying (50)    NOT NULL,
    province    character varying (50)    NOT NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
    
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE domis DROP COLUMN IF EXISTS "gender";
DROP TABLE "domiciles";
-- +goose StatementEnd
