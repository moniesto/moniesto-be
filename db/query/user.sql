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

-- name: UpdateLoginStats :one
UPDATE "user"
SET login_count = login_count + 1,
    last_login = now()
WHERE id = $1
RETURNING *;

-- -- name: GetUserByID :one
-- SELECT (
--         user.id,
--         user.name,
--         user.surname,
--         user.username,
--         user.email,
--         user.email_verified,
--         user.location,
--         user.created_at,
--         user.updated_at,
--         image.link FILTER (WHERE image.type = 1) AS profile_photo_link,
--         image.thumbnail_link FILTER (WHERE image.type = 1) AS profile_photo_thumbnail_link,
--         image.link FILTER (WHERE image.type = 2) AS background_photo_link,
--         image.thumbnail_link FILTER (WHERE image.type = 2) AS background_photo_thumbnail_link,
--     )
-- FROM user
--     INNER JOIN image ON image.user_id = user.id
-- WHERE 
--  user.id = $1;
--  -- name: GetUserByUsername :one
-- SELECT (
--         user.id,
--         user.name,
--         user.surname,
--         user.username,
--         user.email,
--         user.email_verified,
--         user.location,
--         user.created_at,
--         user.updated_at,
--         image.link FILTER (WHERE image.type = 1) AS profile_photo_link,
--         image.thumbnail_link FILTER (WHERE image.type = 1) AS profile_photo_thumbnail_link,
--         image.link FILTER (WHERE image.type = 2) AS background_photo_link,
--         image.thumbnail_link FILTER (WHERE image.type = 2) AS background_photo_thumbnail_link,
--     )
-- FROM user
--     INNER JOIN image ON image.user_id = user.id
-- WHERE 
--  user.username = $1;
-- -- name: GetUserByEmail :one
-- SELECT (
--         user.id,
--         user.name,
--         user.surname,
--         user.username,
--         user.email,
--         user.email_verified,
--         user.location,
--         user.created_at,
--         user.updated_at,
--         image.link FILTER (WHERE image.type = 1) AS profile_photo_link,
--         image.thumbnail_link FILTER (WHERE image.type = 1) AS profile_photo_thumbnail_link,
--         image.link FILTER (WHERE image.type = 2) AS background_photo_link,
--         image.thumbnail_link FILTER (WHERE image.type = 2) AS background_photo_thumbnail_link,
--     )
-- FROM user
--     INNER JOIN image ON image.user_id = user.id
-- WHERE 
--  user.email = $1;
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

-- name: SetPassword :one
UPDATE "user"
SET password = $2
WHERE id = $1
RETURNING *;

-- name: CheckEmail :one
SELECT COUNT(*) = 0 AS isEmailValid
FROM "user"
WHERE email = $1;

-- name: CheckUsername :one
SELECT COUNT(*) = 0 AS isUsernameValid
FROM "user"
WHERE username = $1;