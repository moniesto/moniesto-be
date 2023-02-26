-- name: CreateSubscription :one
INSERT INTO "user_subscription" (
        id,
        user_id,
        moniest_id,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, now(), now())
RETURNING *;

-- name: GetSubscription :one
SELECT *
FROM "user_subscription"
WHERE user_id = $1
    AND moniest_id = $2;

-- name: ActivateSubscription :exec
UPDATE "user_subscription"
SET active = true
WHERE user_id = $1
    AND moniest_id = $2;

-- name: Endsubscription :exec
UPDATE "user_subscription"
SET active = false,
    updated_at = now()
WHERE user_id = $1
    AND moniest_id = $2;

-- name: CheckSubscriptionByMoniestUsername :one
SELECT COUNT(*) != 0 AS subscribed
FROM "user_subscription"
    INNER JOIN "moniest" ON "moniest"."id" = "user_subscription"."moniest_id"
    INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
    AND "user"."username" = $2
WHERE "user_subscription"."active" = TRUE
    AND "user_subscription"."user_id" = $1;