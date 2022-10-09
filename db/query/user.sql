-- name: CreateUser :one
INSERT INTO "user" (
        id,
        name,
        surname,
        username,
        email,
        password
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByID :one
SELECT (
        id,
        name,
        surname,
        username,
        email,
        email_verified,
        location,
        created_at,
        updated_at
    )
FROM "user"
    INNER JOIN "image" ON "image"."user_id" = "user"."id"
WHERE "user"."id" = $1;