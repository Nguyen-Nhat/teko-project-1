-- +goose Up
-- +goose StatementBegin
ALTER TABLE "students"
ADD CONSTRAINT "FK_Student_To_University"
FOREIGN KEY (university_id)
REFERENCES "universities"(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "students"
DROP CONSTRAINT "FK_Student_To_University";
-- +goose StatementEnd
