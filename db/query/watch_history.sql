-- name: AddWatchHistory :one
INSERT INTO watch_history (
    
    video_id,
    watched_at,
    channel_id
)VALUES (
         ?, ?, ?
) RETURNING *;