-- name: GetFeedInfo :one
SELECT * from feeds
WHERE id = $1;
