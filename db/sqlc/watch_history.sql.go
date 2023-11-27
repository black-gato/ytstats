// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: watch_history.sql

package db

import (
	"context"
)

const addWatchHistory = `-- name: AddWatchHistory :one
INSERT INTO watch_history (
    
    video_id,
    watched_at
)VALUES (
         ?, ?
) RETURNING watch_history_id, video_id, watched_at
`

type AddWatchHistoryParams struct {
	VideoID   string
	WatchedAt string
}

func (q *Queries) AddWatchHistory(ctx context.Context, arg AddWatchHistoryParams) (WatchHistory, error) {
	row := q.db.QueryRowContext(ctx, addWatchHistory, arg.VideoID, arg.WatchedAt)
	var i WatchHistory
	err := row.Scan(&i.WatchHistoryID, &i.VideoID, &i.WatchedAt)
	return i, err
}
