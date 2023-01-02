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

-- name: GetPasswordResetTokenByToken :one
SELECT *
FROM "password_reset_token"
WHERE "token" = $1
    AND "deleted" = false;

-- name: DeletePasswordResetTokenByToken :exec
UPDATE "password_reset_token"
SET deleted = true
WHERE "token" = $1;

-- name: DeletePasswordResetTokenByUserID :exec
UPDATE "password_reset_token"
SET deleted = true
WHERE "user_id" = $1;