-- name: AddChannel :one
INSERT INTO channels (
    id,
    channel_name,
    channel_url,
    is_subbed

    ) VALUES (
        ?, ?, ?, ?
) RETURNING *;