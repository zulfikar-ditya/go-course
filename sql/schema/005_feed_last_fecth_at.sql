-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetch_at TIMESTAMP WITH TIME ZONE;

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetch_at;