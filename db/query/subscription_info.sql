-- name: CreateSubscriptionInfo :one
INSERT INTO subscription_info (
        id,
        moniest_id,
        fee,
        message,
        created_at
    )
VALUES ($1, $2, $3, $4, now())
RETURNING *;

-- name: UpdateSubscriptionInfo :one
UPDATE "subscription_info"
SET fee = $2,
    message = $3,
    updated_at = now()
WHERE moniest_id = $1
RETURNING *;

-- name: GetSubscriptionInfoByMoniestId :one
SELECT fee,
    message,
    updated_at
FROM "subscription_info"
WHERE moniest_id = $1;

-- -- name: DeleteSubscriptionInfo :one
-- UPDATE "subscription_info"
-- SET deleted = true,
--     updated_date = now()
-- WHERE moniest_id = $1
-- RETURNING *;
-- -- name: UpdateFee :one
-- UPDATE "subscription_info"
-- SET fee = $2,
--     updated_at = now()
-- WHERE moniest_id = $1
-- RETURNING *;
-- -- name: UpdateMessage :one
-- UPDATE "subscription_info"
-- SET message = $2,
--     updated_at = now()
-- WHERE moniest_id = $1
-- RETURNING *;