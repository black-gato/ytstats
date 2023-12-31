-- name: AddChannel :one
INSERT INTO channels (
    id,
    channel_name,
    channel_url,
    is_subbed

    ) VALUES (
        ?, ?, ?, ?
) RETURNING *;

-- name: GetMostWatchedChannels :many
SELECT channels.channel_name, COUNT(*) AS watch_count, channels.is_subbed
FROM watch_history
INNER JOIN channels ON watch_history.channel_id = channels.id
WHERE channels.is_subbed = 1
GROUP BY watch_history.channel_id
HAVING COUNT(*) >= 1
ORDER BY watch_count DESC
LIMIT :limit;