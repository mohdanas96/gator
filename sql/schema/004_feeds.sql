-- +goose up
ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP;