-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
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
    last_fetched_at = $2,
    updated_at = $3
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
