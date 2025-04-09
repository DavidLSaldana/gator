-- name: GetFeedInfo :one
SELECT * from feeds
WHERE id = $1;

--

-- name: GetFeedID :one
SELECT id from feeds
WHERE url = $1;

--

