-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsByUser :many
SELECT p.*
FROM posts p
JOIN users_feeds uf
ON p.feed_id = uf.feed_id
WHERE uf.user_id = $1;
