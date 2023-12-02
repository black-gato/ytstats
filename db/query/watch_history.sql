-- name: AddWatchHistory :one
INSERT INTO watch_history (
    
    video_id,
    watched_at,
    channel_id
)VALUES (
         ?, ?, ?
) RETURNING *;

-- name: GetMostWatched :many
SELECT videos.video_title, channels.channel_name, COUNT(*) AS watch_count, channels.is_subbed, videos.video_type
FROM watch_history
INNER JOIN videos ON  watch_history.video_id = videos.id
INNER JOIN channels ON watch_history.channel_id = channels.id
WHERE channels.is_subbed = 1
GROUP BY watch_history.video_id
HAVING COUNT(*) > 1
ORDER BY watch_count DESC
LIMIT 10;