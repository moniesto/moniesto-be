-- name: CreateUser :one
INSERT INTO "user" (
        id,
        name,
        surname,
        username,
        email,
        password,
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
    "user"."name",
    "user"."surname",
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
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "subscription_info" ON "subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."id" = $1
    AND "user"."deleted" = false;

-- name: GetUserByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."name",
    "user"."surname",
    "user"."username",
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
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "subscription_info" ON "subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."username" = $1
    AND "user"."deleted" = false;

-- name: GetOwnUserByUsername :one
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
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "subscription_info" ON "subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."username" = $1
    AND "user"."deleted" = false;

-- name: GetUserByEmail :one
SELECT "user"."id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
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
SET name = $2,
    surname = $3,
    location = $4,
    updated_at = now()
WHERE id = $1;

-- name: VerifyEmail :exec
UPDATE "user"
SET email_verified = true,
    updated_at = now()
WHERE id = $1;

-- name: GetActiveUsersVerifiedEmails :many
SELECT email
FROM "user"
WHERE email_verified = true
    AND deleted = false;

-- name: GetInactiveUsersVerifiedEmails :many
SELECT email
FROM "user"
WHERE email_verified = true
    AND deleted = true;

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
    "si"."fee",
    "si"."message",
    "si"."updated_at" as "subscription_info_updated_at",
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
    INNER JOIN "subscription_info" as si ON "si"."moniest_id" = "m"."id"
    AND "us"."user_id" = $1
    AND "us"."active" = TRUE
ORDER BY "us"."created_at" DESC
LIMIT $2 OFFSET $3;