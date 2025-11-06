-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users(username, email, password, role) VALUES(?,?,?,?);

