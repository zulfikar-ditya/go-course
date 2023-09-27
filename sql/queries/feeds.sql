-- name: CreateNewFeed :one
INSERT INTO feeds (id, user_id, url, name, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: GetFeedById :one
SELECT * FROM feeds WHERE id = $1;

-- name: GetFeeds :many
SELECT * FROM feeds ORDER BY created_at DESC;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetch_at DESC NULLS FIRST LIMIT 1;

-- name: UpdateLastFetchFeed :one
UPDATE feeds SET 
last_fetch_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;