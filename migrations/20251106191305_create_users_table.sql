-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
