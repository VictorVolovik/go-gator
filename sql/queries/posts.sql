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
    posts.id,
    posts.created_at,
    posts.updated_at,
    posts.url,
    posts.title,
    posts.description,
    posts.published_at,
    posts.feed_id
FROM posts
INNER JOIN feed_follows
    ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
