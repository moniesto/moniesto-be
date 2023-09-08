-- name: LoginUserByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."password",
    "user"."language",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "moniest_subscription_info"."id" as "moniest_subscription_info_id",
    "moniest_subscription_info"."fee",
    "moniest_subscription_info"."message",
    "moniest_subscription_info"."updated_at" as "moniest_subscription_info_updated_at",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."username" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."username" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."username" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."username" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
    LEFT JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "moniest"."id"
WHERE "user"."username" = $1
    AND "user"."deleted" = false;

-- name: LoginUserByEmail :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."password",
    "user"."language",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "moniest_subscription_info"."id" as "moniest_subscription_info_id",
    "moniest_subscription_info"."fee",
    "moniest_subscription_info"."message",
    "moniest_subscription_info"."updated_at" as "moniest_subscription_info_updated_at",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
    LEFT JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "moniest"."id"
WHERE "user"."email" = $1
    AND "user"."deleted" = false;

-- name: UpdateLoginStats :one
UPDATE "user"
SET login_count = login_count + 1,
    last_login = now(),
    updated_at = now()
WHERE id = $1
RETURNING *;