-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE genres (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE genres;
-- +goose StatementEnd
