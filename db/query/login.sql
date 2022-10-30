-- name: LoginUserByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."password",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
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
WHERE "user"."username" = $1;

-- name: LoginUserByEmail :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."password",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = "profile_photo"
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = "profile_photo"
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = "background_photo"
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
                INNER JOIN "user" ON "user"."id" = "image"."user_id"
            WHERE "user"."email" = $1
                AND "image"."type" = "background_photo"
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE "user"."email" = $1;