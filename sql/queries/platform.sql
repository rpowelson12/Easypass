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

-- name: GetPassword :one
SELECT password FROM platform
WHERE user_id = $1 
AND platform = $2;

-- name: GetPlatforms :many
SELECT * FROM platform
WHERE user_id = $1;

-- name: DeletePlatform :exec
DELETE FROM platform
WHERE platform = $1;

-- name: UpdatePassword :exec
UPDATE platform
SET password=$1
WHERE user_id=$2 AND platform=$3;

-- name: GetPlatform :one
SELECT * From platform
WHERE platform = $1;

-- name: AddPassword :exec
INSERT INTO platform (id, created_at, updated_at, platform, password, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);
