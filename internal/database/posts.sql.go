// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
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
RETURNING id, created_at, updated_at, url, title, description, published_at, feed_id
`

type CreatePostParams struct {
	Url         string
	Title       string
	Description string
	PublishedAt string
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Url,
		arg.Title,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
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
LIMIT $2
`

type GetPostsParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPosts, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Url,
			&i.Title,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
