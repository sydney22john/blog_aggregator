-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    url TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE feeds;
