-- name:GetNextFeedToFetch :one
SELECT * FROM feeds
WHERE Min(last_fetched_at) = last_fetched_at;
--
