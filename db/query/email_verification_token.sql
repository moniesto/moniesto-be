-- name: CreateEmailVerificationToken :one
INSERT INTO "email_verification_token" (
        id,
        user_id,
        token,
        token_expiry,
        redirect_url,
        deleted,
        created_at
    )
VALUEs ($1, $2, $3, $4, $5, false, now())
RETURNING *;

-- name: GetEmailVerificationTokenByToken :one
SELECT *
FROM "email_verification_token"
WHERE "token" = $1
    AND "deleted" = false;

-- name: DeleteEmailVerificationTokenByToken :exec
UPDATE "email_verification_token"
SET deleted = true
WHERE "token" = $1;

-- name: DeleteEmailVerificationTokenByUserID :exec
UPDATE "email_verification_token"
SET deleted = true
WHERE "user_id" = $1;