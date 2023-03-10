-- name: CreateMoniest :one
INSERT INTO moniest (
        id,
        user_id,
        bio,
        description,
        created_at
    )
VALUES ($1, $2, $3, $4, now())
RETURNING *;

-- -- name: DeleteMoniest :one
-- UPDATE moniest
-- SET deleted = true,
--     updated_at = now()
-- WHERE moniest.id = $1
-- RETURNING *;
-- name: UpdateMoniest :one
UPDATE moniest
SET bio = $2,
    description = $3,
    updated_at = now()
WHERE moniest.id = $1
RETURNING *;

-- name: GetMoniestByUserId :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    "subscription_info"."id" as "subscription_info_id",
    "subscription_info"."fee",
    "subscription_info"."message",
    "subscription_info"."updated_at" as "subscription_info_updated_at",
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
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    INNER JOIN "subscription_info" ON "subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."id" = $1
    AND "user"."deleted" = false;

-- name: GetMoniestByMoniestId :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    "subscription_info"."id" as "subscription_info_id",
    "subscription_info"."fee",
    "subscription_info"."message",
    "subscription_info"."updated_at" as "subscription_info_updated_at",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
            WHERE "moniest"."id" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
            WHERE "moniest"."id" = $1
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
            WHERE "moniest"."id" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
            WHERE "moniest"."id" = $1
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    INNER JOIN "subscription_info" ON "subscription_info"."moniest_id" = "moniest"."id"
WHERE "moniest"."id" = $1
    AND "user"."deleted" = false;

-- name: GetMoniestByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    "subscription_info"."id" as "subscription_info_id",
    "subscription_info"."fee",
    "subscription_info"."message",
    "subscription_info"."updated_at" as "subscription_info_updated_at",
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
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    INNER JOIN "subscription_info" ON "subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."username" = $1
    AND "user"."deleted" = false;

-- name: CheckUserIsMoniestByUsername :one
SELECT COUNT(*) != 0 AS userIsMoniest
FROM "moniest"
    INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
    AND "user"."username" = $1;

-- name: CheckUserIsMoniestByID :one
SELECT COUNT(*) != 0 AS userIsMoniest
FROM "moniest"
WHERE "moniest"."user_id" = $1;

-- name: GetMoniestStatsByUsername :one
SELECT COUNT(DISTINCT "us1"."id") as "subscription_count",
    COUNT(DISTINCT "us2"."id") as "subscriber_count",
    COUNT(DISTINCT "pc"."id") as "post_count"
FROM "user"
    LEFT JOIN "user_subscription" as "us1" ON "us1"."user_id" = "user"."id"
    AND "us1"."active" = TRUE
    LEFT JOIN "moniest" as "m" ON "m"."user_id" = "user"."id"
    LEFT JOIN "user_subscription" as "us2" ON "us2"."moniest_id" = "m"."id"
    AND "us2"."active" = TRUE
    LEFT JOIN "post_crypto" as "pc" ON "pc"."moniest_id" = "m"."id"
where "user"."username" = $1;

-- -- name: GetMoniestByEmail :one
-- SELECT "user"."id",
--     "moniest"."id" as "moniest_id",
--     "user"."name",
--     "user"."surname",
--     "user"."username",
--     "user"."email",
--     "user"."email_verified",
--     "user"."location",
--     "user"."created_at",
--     "user"."updated_at",
--     "moniest"."bio",
--     "moniest"."description",
--     "moniest"."score",
--     COALESCE (
--         (
--             SELECT "image"."link"
--             FROM "image"
--                 INNER JOIN "user" ON "user"."id" = "image"."user_id"
--             WHERE "user"."email" = $1
--                 AND "image"."type" = 'profile_photo'
--         ),
--         ''
--     ) AS "profile_photo_link",
--     COALESCE (
--         (
--             SELECT "image"."thumbnail_link"
--             FROM "image"
--                 INNER JOIN "user" ON "user"."id" = "image"."user_id"
--             WHERE "user"."email" = $1
--                 AND "image"."type" = 'profile_photo'
--         ),
--         ''
--     ) AS "profile_photo_thumbnail_link",
--     COALESCE (
--         (
--             SELECT "image"."link"
--             FROM "image"
--                 INNER JOIN "user" ON "user"."id" = "image"."user_id"
--             WHERE "user"."email" = $1
--                 AND "image"."type" = 'background_photo'
--         ),
--         ''
--     ) AS "background_photo_link",
--     COALESCE (
--         (
--             SELECT "image"."thumbnail_link"
--             FROM "image"
--                 INNER JOIN "user" ON "user"."id" = "image"."user_id"
--             WHERE "user"."email" = $1
--                 AND "image"."type" = 'background_photo'
--         ),
--         ''
--     ) AS "background_photo_thumbnail_link"
-- FROM "user"
--     INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
-- WHERE "user"."email" = $1
--     AND "user"."deleted" = false;