-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1
)
RETURNING *;

-- name: GetUserByName :one
SELECT
    id,
    created_at,
    updated_at,
    name
FROM users
WHERE name = $1;

-- name: GetAllUsers :many
SELECT
    id,
    created_at,
    updated_at,
    name
FROM users;

-- name: DeleteUsers :exec
DELETE FROM users;
