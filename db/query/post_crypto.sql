-- name: CreatePost :one
INSERT INTO "post_crypto" (
        id,
        moniest_id,
        base_currency,
        quote_currency,
        duration,
        target1,
        target2,
        target3,
        stop,
        direction,
        score,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        now(),
        now()
    )
RETURNING *;

-- name: DeletePost :one
UPDATE "post_crypto"
SET "deleted" = true,
    "updated_at" = now()
WHERE "id" = $1
RETURNING *;

-- TODO:
-- Post update queryleri yazılacak! her field overwrite edilecek mi?
-- FIXME:
-- tahmin girildikten sonra direction değiştirilebilir mi?
-- name: UpdatePost :one
UPDATE "post_crypto"
SET "duration" = $2,
    "target1" = $3,
    "target2" = $4,
    "target3" = $5,
    "stop" = $6,
    "updated_at" = now()
WHERE "id" = $1
RETURNING *;