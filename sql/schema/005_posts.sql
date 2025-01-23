-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    url TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    published_at TEXT NOT NULL,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_posts_feed_id
    FOREIGN KEY (feed_id)
    REFERENCES feeds (id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
