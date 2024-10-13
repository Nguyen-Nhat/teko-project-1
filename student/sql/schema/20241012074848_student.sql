-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Student" (
    "id" integer PRIMARY KEY,
    "fullname" varchar NOT NULL DEFAULT '',
    "sex" smallint NOT NULL DEFAULT '0',
    "dob" date NOT NULL,
    "university_id" integer NOT NULL,
    "enrollment_year" integer NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Student";
-- +goose StatementEnd