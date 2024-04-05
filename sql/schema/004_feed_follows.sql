-- +goose Up
CREATE TABLE users_feeds (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users_feeds;
