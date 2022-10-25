-- name: AddPostDescription :one
INSERT INTO "post_crypto_description" (
    id,
    post_id,
    description,
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, now(), now()
)
RETURNING *;

-- name: UpdateDescription :one
UPDATE "post_crypto_description" 
SET "description" = $2,
    "updated_at" = now()
WHERE post_id = $1
RETURNING *;
