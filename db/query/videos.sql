-- name: AddVideo :one
INSERT INTO videos (
    id,
    video_type,
    video_title,
    channel_id
)VALUES (
        ?, ?, ?, ?
) RETURNING *;