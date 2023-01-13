// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: image.sql

package db

import (
	"context"
)

const addImage = `-- name: AddImage :one
INSERT INTO image (
        id,
        user_id,
        link,
        thumbnail_link,
        type,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, now())
RETURNING id, user_id, link, thumbnail_link, type, created_at, updated_at
`

type AddImageParams struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Link          string    `json:"link"`
	ThumbnailLink string    `json:"thumbnail_link"`
	Type          ImageType `json:"type"`
}

func (q *Queries) AddImage(ctx context.Context, arg AddImageParams) (Image, error) {
	row := q.db.QueryRowContext(ctx, addImage,
		arg.ID,
		arg.UserID,
		arg.Link,
		arg.ThumbnailLink,
		arg.Type,
	)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Link,
		&i.ThumbnailLink,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBackgroundPhoto = `-- name: GetBackgroundPhoto :one
SELECT image.link,
    image.thumbnail_link
FROM image
WHERE image.user_id = $1
    and image.type = 'background_photo'
`

type GetBackgroundPhotoRow struct {
	Link          string `json:"link"`
	ThumbnailLink string `json:"thumbnail_link"`
}

func (q *Queries) GetBackgroundPhoto(ctx context.Context, userID string) (GetBackgroundPhotoRow, error) {
	row := q.db.QueryRowContext(ctx, getBackgroundPhoto, userID)
	var i GetBackgroundPhotoRow
	err := row.Scan(&i.Link, &i.ThumbnailLink)
	return i, err
}

const getImagesByUserId = `-- name: GetImagesByUserId :many
SELECT image.link,
    image.thumbnail_link,
    image.type
FROM image
WHERE image.user_id = $1
`

type GetImagesByUserIdRow struct {
	Link          string    `json:"link"`
	ThumbnailLink string    `json:"thumbnail_link"`
	Type          ImageType `json:"type"`
}

func (q *Queries) GetImagesByUserId(ctx context.Context, userID string) ([]GetImagesByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getImagesByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetImagesByUserIdRow{}
	for rows.Next() {
		var i GetImagesByUserIdRow
		if err := rows.Scan(&i.Link, &i.ThumbnailLink, &i.Type); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfilePhoto = `-- name: GetProfilePhoto :one
SELECT link,
    thumbnail_link
FROM image
WHERE user_id = $1
    and type = 'profile_photo'
`

type GetProfilePhotoRow struct {
	Link          string `json:"link"`
	ThumbnailLink string `json:"thumbnail_link"`
}

func (q *Queries) GetProfilePhoto(ctx context.Context, userID string) (GetProfilePhotoRow, error) {
	row := q.db.QueryRowContext(ctx, getProfilePhoto, userID)
	var i GetProfilePhotoRow
	err := row.Scan(&i.Link, &i.ThumbnailLink)
	return i, err
}

const updateBackgroundPhoto = `-- name: UpdateBackgroundPhoto :one
UPDATE image
SET link = $2,
    thumbnail_link = $3
WHERE image.user_id = $1
    and image.type = 'background_photo'
RETURNING id, user_id, link, thumbnail_link, type, created_at, updated_at
`

type UpdateBackgroundPhotoParams struct {
	UserID        string `json:"user_id"`
	Link          string `json:"link"`
	ThumbnailLink string `json:"thumbnail_link"`
}

func (q *Queries) UpdateBackgroundPhoto(ctx context.Context, arg UpdateBackgroundPhotoParams) (Image, error) {
	row := q.db.QueryRowContext(ctx, updateBackgroundPhoto, arg.UserID, arg.Link, arg.ThumbnailLink)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Link,
		&i.ThumbnailLink,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProfilePhoto = `-- name: UpdateProfilePhoto :one
UPDATE image
SET link = $2,
    thumbnail_link = $3
WHERE image.user_id = $1
    and image.type = 'profile_photo'
RETURNING id, user_id, link, thumbnail_link, type, created_at, updated_at
`

type UpdateProfilePhotoParams struct {
	UserID        string `json:"user_id"`
	Link          string `json:"link"`
	ThumbnailLink string `json:"thumbnail_link"`
}

func (q *Queries) UpdateProfilePhoto(ctx context.Context, arg UpdateProfilePhotoParams) (Image, error) {
	row := q.db.QueryRowContext(ctx, updateProfilePhoto, arg.UserID, arg.Link, arg.ThumbnailLink)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Link,
		&i.ThumbnailLink,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
