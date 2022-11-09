-- name: CreatePasswordResetToken :one
INSERT INTO "password_reset_token" (
        id,
        user_id,
        token,
        token_expiry,
        deleted,
        created_at
    )
VALUEs ($1, $2, $3, $4, false, now())
RETURNING *;

-- -- name: GetPasswordResetTokenByToken :one
-- SELECT * FROM