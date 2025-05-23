// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: markFeedFetched.sql

package database

import (
	"context"
	"database/sql"
)

const markFeedFetched = `-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2, updated_at = $2
WHERE id = $1
`

type MarkFeedFetchedParams struct {
	ID            int32
	LastFetchedAt sql.NullTime
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.ID, arg.LastFetchedAt)
	return err
}
