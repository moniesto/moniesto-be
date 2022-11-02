-- name: AddImage :one
INSERT INTO image (
    id, 
    user_id, 
    link, 
    thumbnail_link, 
    type,
    created_at) 
VALUES ($1, $2, $3, $4, $5, now())
RETURNING *;

-- name: GetImagesByUserId :many
SELECT image.link, image.thumbnail_link, image.type
FROM image
WHERE image.user_id = $1;


-- name: GetProfilePhoto :one
SELECT link, thumbnail_link
FROM image
WHERE user_id = $1 and type = 'profile_photo';

-- name: GetBackgroundPhoto :one
SELECT image.link, image.thumbnail_link
FROM image
WHERE image.user_id = $1 and image.type = 'background_photo';


-- name: UpdateProfilePhoto :one
UPDATE image
SET link = $2,
    thumbnail_link = $3
WHERE image.user_id = $1 and image.type = 'profile_photo'
RETURNING *;

-- name: UpdateBackgroundPhoto :one
UPDATE image
SET link = $2,
    thumbnail_link = $3
WHERE image.user_id = $1 and image.type = 'background_photo'
RETURNING *;

-- -- delete pp 

-- --delete background