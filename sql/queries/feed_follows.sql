-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT
    iff.id,
    iff.created_at,
    iff.updated_at,
    iff.user_id,
    iff.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM inserted_feed_follow AS iff
INNER JOIN users
    ON iff.user_id = users.id
INNER JOIN feeds
    ON iff.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id,
    ff.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follows AS ff
INNER JOIN users
    ON ff.user_id = users.id
INNER JOIN feeds
    ON ff.feed_id = feeds.id
WHERE ff.user_id = $1;
