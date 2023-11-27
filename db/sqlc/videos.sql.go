// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: videos.sql

package db

import (
	"context"
)

const addVideo = `-- name: AddVideo :one
INSERT INTO videos (
    id,
    video_type,
    video_title,
    channel_id
)VALUES (
        ?, ?, ?, ?
) RETURNING id, video_type, video_title, channel_id
`

type AddVideoParams struct {
	ID         string
	VideoType  string
	VideoTitle string
	ChannelID  string
}

func (q *Queries) AddVideo(ctx context.Context, arg AddVideoParams) (Video, error) {
	row := q.db.QueryRowContext(ctx, addVideo,
		arg.ID,
		arg.VideoType,
		arg.VideoTitle,
		arg.ChannelID,
	)
	var i Video
	err := row.Scan(
		&i.ID,
		&i.VideoType,
		&i.VideoTitle,
		&i.ChannelID,
	)
	return i, err
}

const countVideo = `-- name: CountVideo :many

SELECT video_id, watch_history_id, videos.video_title, COUNT(*) AS watch_count
FROM watch_history
INNER JOIN videos ON  watch_history.video_id = videos.id
GROUP BY video_id, watch_history_id, videos.video_title
HAVING COUNT(*) > 1
ORDER BY watch_count DESC
`

type CountVideoRow struct {
	VideoID        string
	WatchHistoryID int64
	VideoTitle     string
	WatchCount     int64
}

func (q *Queries) CountVideo(ctx context.Context) ([]CountVideoRow, error) {
	rows, err := q.db.QueryContext(ctx, countVideo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CountVideoRow
	for rows.Next() {
		var i CountVideoRow
		if err := rows.Scan(
			&i.VideoID,
			&i.WatchHistoryID,
			&i.VideoTitle,
			&i.WatchCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
