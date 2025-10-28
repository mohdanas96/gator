-- name: CreateFeedFollow :many
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT inserted_feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM inserted_feed_follows
JOIN feeds on inserted_feed_follows.feed_id = feeds.id
JOIN users on inserted_feed_follows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name as feed_name, users.name as user_name 
FROM feed_follows
JOIN feeds on feed_follows.feed_id = feeds.id
JOIN users on feed_follows.user_id = users.id
WHERE users.id = $1;