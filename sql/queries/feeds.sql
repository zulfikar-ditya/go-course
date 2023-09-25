-- name: CreateNewFeed :one
INSERT INTO feeds (id, user_id, url, name, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING id, user_id, url, name, created_at, updated_at;

-- name: GetFeedById :one
SELECT id, user_id, url, name, created_at, updated_at FROM feeds WHERE id = $1;

-- name: GetFeeds :many
SELECT id, user_id, url, name, created_at, updated_at FROM feeds ORDER BY created_at DESC;

-- name: GetFeedByUserId
/* SELECT id, user_id, url, name, created_at, updated_at FROM feeds WHERE user_id = $1;

-- name: GetFeedByUrl
SELECT id, user_id, url, name, created_at, updated_at FROM feeds WHERE url = $1; */