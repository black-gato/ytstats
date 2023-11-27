-- name: AddVideo :one
INSERT INTO videos (
    id,
    video_type,
    video_title,
    channel_id
)VALUES (
        ?, ?, ?, ?
) RETURNING *;

-- name: CountVideo :many

SELECT video_id, watch_history_id, videos.video_title, COUNT(*) AS watch_count
FROM watch_history
INNER JOIN videos ON  watch_history.video_id = videos.id
GROUP BY video_id, watch_history_id, videos.video_title
HAVING COUNT(*) > 1
ORDER BY watch_count DESC;
