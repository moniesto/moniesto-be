-- name: CreateSubscription :one
INSERT INTO "user_subscription" (
    id,
    user_id,
    moniest_id,
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, now(), now()
)
RETURNING *;

-- name: Endsubscription :one
UPDATE "user_subscription" 
SET "deleted" = true,
    "updated_at" = now()
WHERE "user_id" = $1 AND "moniest_id" = $2
RETURNING *;

-- TODO:
    -- delete field'ını update ettiğimiz için resub durumunu handle etmemiz gerek
