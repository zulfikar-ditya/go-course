-- name: CreateUser :one
INSERT INTO user (id, name, created_at, updated_at) 
VALUES ($1, $2, $3, $4)
RETURNING id, name, created_at, updated_at;