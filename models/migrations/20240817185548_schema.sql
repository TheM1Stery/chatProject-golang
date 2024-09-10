-- +goose Up
-- +goose StatementBegin
CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    name TEXT,
    pass TEXT

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE test
-- +goose StatementEnd
