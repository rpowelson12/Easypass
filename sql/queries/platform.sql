-- name: GeneratePassword :one
INSERT INTO platform (id, created_at, updated_at, platform, password, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

