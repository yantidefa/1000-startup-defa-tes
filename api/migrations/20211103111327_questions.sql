-- +goose Up
-- +goose StatementBegin

CREATE TABLE "questions" (
    id          serial PRIMARY KEY        NOT NULL,
    questions        character varying    NOT NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
    
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE domis DROP COLUMN IF EXISTS "gender";
DROP TABLE "questions";
-- +goose StatementEnd
