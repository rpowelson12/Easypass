-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, password)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserId :one
SELECT id FROM users WHERE name = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE name = $1;
