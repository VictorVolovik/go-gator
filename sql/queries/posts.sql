-- name: CreatePost :one
INSERT INTO posts (
    id, created_at, updated_at, url, title, description, published_at, feed_id
)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetPosts :many
SELECT
    id,
    created_at,
    updated_at,
    url,
    title,
    description,
    published_at,
    feed_id
FROM posts
ORDER BY published_at DESC
LIMIT $1;
