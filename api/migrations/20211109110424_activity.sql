-- +goose Up
-- +goose StatementBegin
CREATE TABLE "activities" (
    id          serial PRIMARY KEY        NOT NULL,
    name        character varying(50)      NOT NULL,
    description       text             NOT NULL,
    image       character varying            NULL,
    location       text             NOT NULL,
    start       timestamp         NOT NULL,
    finish       timestamp         NOT NULL,
     created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "activities";
-- +goose StatementEnd
