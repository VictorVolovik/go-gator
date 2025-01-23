-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT
    id,
    created_at,
    updated_at,
    name,
    url
FROM feeds
WHERE url = $1;

-- name: GetAllFeeds :many
SELECT
    feeds.id,
    feeds.created_at,
    feeds.updated_at,
    feeds.name,
    feeds.url,
    feeds.user_id,
    users.name AS user_name
FROM feeds
INNER JOIN users
    ON feeds.user_id = users.id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET
    last_fetched_at = now(),
    updated_at = now()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT
    id,
    created_at,
    updated_at,
    name,
    url
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
