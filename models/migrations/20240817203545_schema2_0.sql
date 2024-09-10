-- +goose Up
-- +goose StatementBegin
ALTER TABLE test ADD CONSTRAINT ux_test_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE test DROP CONSTRAINT ux_test_name;
-- +goose StatementEnd
