-- name: CreateMoniest :one
INSERT INTO "moniest" (
        id,
        user_id,
        bio,
        description
    )
VALUES ($1, $2, $3, $4)
RETURNING *;