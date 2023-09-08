-- name: CreateUser :one
INSERT INTO "user" (
        id,
        fullname,
        username,
        email,
        password,
        language,
        created_at,
        last_login
    )
VALUES ($1, $2, $3, $4, $5, $6, now(), now())
RETURNING *;

-- name: DeleteUser :one
UPDATE "user"
SET deleted = true,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: GetUserByID :one
SELECT "user"."id",
    "user"."fullname",
    "user"."username",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
WHERE "user"."id" = $1
    AND "user"."deleted" = false;

-- name: GetOwnUserByID :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
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
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
    LEFT JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "moniest"."id"
WHERE "user"."id" = $1
    AND "user"."deleted" = false;

-- name: GetUserByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email_verified",
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

-- name: GetOwnUserByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
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

-- name: GetUserByEmail :one
SELECT "user"."id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."language",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
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
WHERE "user"."email" = $1
    AND "user"."deleted" = false;

-- name: UpdateUser :exec
UPDATE "user"
SET fullname = $2,
    location = $3,
    language = $4,
    updated_at = now()
WHERE id = $1;

-- name: VerifyEmail :exec
UPDATE "user"
SET email_verified = true,
    updated_at = now()
WHERE id = $1;

-- name: GetPasswordByID :one
SELECT password
FROM "user"
WHERE id = $1
    AND "user"."deleted" = false;

-- name: SetPassword :exec
UPDATE "user"
SET password = $2,
    updated_at = now()
WHERE id = $1;

-- name: SetUsername :exec
UPDATE "user"
SET username = $2,
    updated_at = now()
WHERE id = $1;

-- name: CheckEmail :one
SELECT COUNT(*) = 0 AS isEmailValid
FROM "user"
WHERE email = $1;

-- name: CheckUsername :one
SELECT COUNT(*) = 0 AS isUsernameValid
FROM "user"
WHERE username = $1;

-- name: GetSubscriptions :many
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
FROM "moniest" as m
    INNER JOIN "user_subscription" AS us ON "m"."id" = "us"."moniest_id"
    INNER JOIN "user" as u ON "u"."id" = "m"."user_id"
    INNER JOIN "moniest_subscription_info" as msi ON "msi"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    AND "us"."user_id" = $1
    AND "us"."active" = TRUE
ORDER BY "us"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetUserStatsByUsername :one
SELECT COUNT("us"."id") as "user_subscription_count"
FROM "user"
    LEFT JOIN "user_subscription" as "us" ON "us"."user_id" = "user"."id"
    AND "us"."active" = TRUE
where "user"."username" = $1;

-- name: GetUserLanguageByEmail :one
SELECT "language"
FROM "user"
WHERE "email" = $1;

-- -- name: GetUserLanguageByID :one
-- SELECT "language"
-- FROM "user"
-- WHERE "id" = $1;
-- -- name: GetUserLanguageByUsername :one
-- SELECT "language"
-- FROM "user"
-- WHERE "username" = $1;