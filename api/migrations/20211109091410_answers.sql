-- +goose Up
-- +goose StatementBegin
CREATE TABLE "answers" (
    id          serial PRIMARY KEY        NOT NULL,
    answer       character varying             NOT NULL,
    point       int             NOT NULL,
    id_question       int             NOT NULL,
    id_roles_point       int             NOT NULL,
    created_at timestamp              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL
);
-- +goose StatementEnd
ALTER TABLE "answers" ADD FOREIGN KEY ("id_question") REFERENCES "questions" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE "answers" ADD FOREIGN KEY ("id_roles_point") REFERENCES "roles_points" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT;
-- +goose Down
-- +goose StatementBegin
DROP TABLE "answers";
-- +goose StatementEnd
