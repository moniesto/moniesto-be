-- name: CreateSubscription :one
INSERT INTO "user_subscription" (
        id,
        user_id,
        moniest_id,
        latest_transaction_id,
        subscription_start_date,
        subscription_end_date,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, now(), now())
RETURNING *;

-- name: GetSubscription :one
SELECT *
FROM "user_subscription"
WHERE user_id = $1
    AND moniest_id = $2;

-- name: ActivateSubscription :exec
UPDATE "user_subscription"
SET active = true,
    latest_transaction_id = $3,
    subscription_start_date = $4,
    subscription_end_date = $5,
    updated_at = now()
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

-- name: GetUserSubscriptionInfo :one
SELECT "user"."id" as "user_id",
    "moniest"."id" as "moniest_id",
    "user_subscription"."latest_transaction_id" as "transaction_id",
    "user_subscription"."subscription_start_date",
    "user_subscription"."subscription_end_date",
    "binance_payout_history"."payer_id",
    "binance_payout_history"."amount"
FROM "user_subscription"
    INNER JOIN "moniest" ON "moniest"."id" = "user_subscription"."moniest_id"
    INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
    INNER JOIN "binance_payout_history" ON "binance_payout_history"."transaction_id" = "user_subscription"."latest_transaction_id"
    AND "binance_payout_history"."date_index" = 1
    AND "user"."username" = $2
WHERE "user_subscription"."active" = TRUE
    AND "user_subscription"."user_id" = $1;

-- name: GetSubscribers :many
SELECT "u"."id",
    "m"."id" as "moniest_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "u"."created_at",
    "u"."updated_at",
    "m"."bio",
    "m"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "msi"."id" as "moniest_subscription_info_id",
    "msi"."fee",
    "msi"."message",
    "msi"."updated_at" as "moniest_subscription_info_updated_at",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user" as u
    INNER JOIN "user_subscription" as us ON "us"."user_id" = "u"."id"
    AND "us"."active" = TRUE
    LEFT JOIN "moniest" as m ON "m"."user_id" = "u"."id"
    LEFT JOIN "moniest_subscription_info" as msi ON "msi"."moniest_id" = "m"."id"
    LEFT JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
WHERE "us"."moniest_id" = $1
LIMIT $2 OFFSET $3;

-- name: GetSubscribersBriefs :many
SELECT "u"."id",
    "u"."fullname",
    "u"."username",
    "u"."email",
    "u"."language"
FROM "user" as u
    INNER JOIN "user_subscription" as us ON "us"."user_id" = "u"."id"
    AND "us"."active" = TRUE
WHERE "us"."moniest_id" = $1;