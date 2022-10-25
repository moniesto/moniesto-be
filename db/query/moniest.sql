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

-- name: DeleteMoniest :one
UPDATE moniest 
SET deleted = true,
    updated_at = now()
WHERE moniest.id = $1
RETURNING *;

-- name: UpdateMoniestBio :one
UPDATE moniest
SET bio = $2,
    updated_at = now()
WHERE moniest.id = $1
RETURNING *;

-- name: UpdateMoniestDescription :one
UPDATE moniest
SET description = $2,
    updated_at = now()
WHERE moniest.id = $1
RETURNING *;

-- name: GetMoniestByUserId :one
SELECT "user"."id",
        "moniest"."id",
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
        (SELECT "image"."link" 
            FROM "image" 
            WHERE "image"."user_id" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_link",
        (SELECT "image"."thumbnail_link" 
            FROM "image"
            WHERE "image"."user_id" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_thumbnail_link",
        (SELECT "image"."link" 
            FROM "image" 
            WHERE "image"."user_id" = $1 
            AND "image"."type" = "background_photo") 
        AS "background_photo_link",
        (SELECT "image"."thumbnail_link" 
            FROM "image" 
            WHERE "image"."user_id" = $1 
            AND "image"."type" = "background_photo") 
        AS "background_photo_thumbnail_link"
FROM "user" 
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE 
    "user"."id" = $1;



-- name: GetMoniestByMoniestId :one
SELECT "user"."id",
        "moniest"."id",
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
        (SELECT "image"."link" 
        FROM "image" 
            INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
        WHERE "moniest"."id" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_link",

        (SELECT "image"."thumbnail_link" 
        FROM "image"
            INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
        WHERE "moniest"."id" = $1
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_thumbnail_link",
        
        (SELECT "image"."link" 
        FROM "image" 
            INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
        WHERE "moniest"."id" = $1
            AND "image"."type" = "background_photo") 
        AS "background_photo_link",

        (SELECT "image"."thumbnail_link" 
        FROM "image" 
            INNER JOIN "moniest" ON "moniest"."user_id" = "image"."user_id"
        WHERE "moniest"."id" = $1
            AND "image"."type" = "background_photo") 
        AS "background_photo_thumbnail_link"
FROM "user"
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE 
    "moniest"."id" = $1;

-- name: GetMoniestByEmail :one
SELECT "user"."id",
        "moniest"."id",
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
        (SELECT "image"."link" 
        FROM "image" 
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_link",

        (SELECT "image"."thumbnail_link" 
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_thumbnail_link",
        
        (SELECT "image"."link" 
        FROM "image" 
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1 
            AND "image"."type" = "background_photo") 
        AS "background_photo_link",

        (SELECT "image"."thumbnail_link" 
        FROM "image" 
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1 
            AND "image"."type" = "background_photo") 
        AS "background_photo_thumbnail_link"

FROM "user"
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE 
    "user"."email" = $1;

-- name: GetMoniestByUsername :one
SELECT "user"."id",
        "moniest"."id",
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
        (SELECT "image"."link" 
        FROM "image" 
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_link",

        (SELECT "image"."thumbnail_link" 
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1 
            AND "image"."type" = "profile_photo") 
        AS "profile_photo_thumbnail_link",
        
        (SELECT "image"."link" 
        FROM "image" 
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1 
            AND "image"."type" = "background_photo") 
        AS "background_photo_link",

        (SELECT "image"."thumbnail_link" 
        FROM "image" 
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1 
            AND "image"."type" = "background_photo") 
        AS "background_photo_thumbnail_link"
FROM "user"
    INNER JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE 
    "user"."username" = $1;