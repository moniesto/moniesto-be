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
SET active = true,
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

-- name: GetSubscribers :many
SELECT "u"."id",
    "m"."id" as "moniest_id",
    "u"."name",
    "u"."surname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "u"."created_at",
    "u"."updated_at",
    "m"."bio",
    "m"."description",
    "m"."score",
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
WHERE "us"."moniest_id" = $1
LIMIT $2 OFFSET $3;