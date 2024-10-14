-- +goose Up
-- +goose StatementBegin
ALTER TABLE "Student"
ADD CONSTRAINT "FK_Student_To_University"
FOREIGN KEY (university_id)
REFERENCES "University"(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "Student"
DROP CONSTRAINT "FK_Student_To_University";
-- +goose StatementEnd
