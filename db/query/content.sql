-- name: GetSubscribedActivePosts :many
SELECT "pc"."id",
    "pc"."currency",
    "pc"."start_price",
    "pc"."duration",
    "pc"."target1",
    "pc"."target2",
    "pc"."target3",
    "pc"."stop",
    "pc"."direction",
    "pc"."finished",
    "pc"."status",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "m"."score" as "moniest_score",
    "u"."id" as "user_id",
    "u"."name",
    "u"."surname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
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
FROM "post_crypto" AS pc
    INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
    AND "us"."user_id" = $1
    AND "pc"."duration" > now()
    AND "pc"."finished" = FALSE
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetSubscribedDeactivePosts :many
SELECT "pc"."id",
    "pc"."currency",
    "pc"."start_price",
    "pc"."duration",
    "pc"."target1",
    "pc"."target2",
    "pc"."target3",
    "pc"."stop",
    "pc"."direction",
    "pc"."finished",
    "pc"."status",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "m"."score" as "moniest_score",
    "u"."id" as "user_id",
    "u"."name",
    "u"."surname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
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
FROM "post_crypto" AS pc
    INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
    AND "us"."user_id" = $1
    AND (
        "pc"."duration" < now()
        OR "pc"."finished" = TRUE
    )
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetDeactivePosts :many
SELECT "pc"."id",
    "pc"."currency",
    "pc"."start_price",
    "pc"."duration",
    "pc"."target1",
    "pc"."target2",
    "pc"."target3",
    "pc"."stop",
    "pc"."direction",
    "pc"."finished",
    "pc"."status",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "m"."score" as "moniest_score",
    "u"."id" as "user_id",
    "u"."name",
    "u"."surname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
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
FROM "post_crypto" AS pc
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    AND (
        "pc"."duration" < now()
        OR "pc"."finished" = TRUE
    )
    AND "pc"."status" = 'success'
ORDER BY "pc"."score" DESC
LIMIT $1 OFFSET $2;