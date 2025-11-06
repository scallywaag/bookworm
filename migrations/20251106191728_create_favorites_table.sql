-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE favorites (
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, book_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE favorites;
-- +goose StatementEnd
