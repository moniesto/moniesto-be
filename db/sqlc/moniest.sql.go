// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: moniest.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMoniest = `-- name: CreateMoniest :one
INSERT INTO moniest (
        id,
        user_id,
        bio,
        description,
        created_at
    )
VALUES ($1, $2, $3, $4, now())
RETURNING id, user_id, bio, description, score, deleted, created_at, updated_at
`

type CreateMoniestParams struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Bio         sql.NullString `json:"bio"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateMoniest(ctx context.Context, arg CreateMoniestParams) (Moniest, error) {
	row := q.db.QueryRowContext(ctx, createMoniest,
		arg.ID,
		arg.UserID,
		arg.Bio,
		arg.Description,
	)
	var i Moniest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMoniestByMoniestId = `-- name: GetMoniestByMoniestId :one
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
    AND "user"."deleted" = false
`

type GetMoniestByMoniestIdRow struct {
	ID                           string         `json:"id"`
	MoniestID                    string         `json:"moniest_id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	Bio                          sql.NullString `json:"bio"`
	Description                  sql.NullString `json:"description"`
	Score                        float64        `json:"score"`
	SubscriptionInfoID           string         `json:"subscription_info_id"`
	Fee                          float64        `json:"fee"`
	Message                      sql.NullString `json:"message"`
	SubscriptionInfoUpdatedAt    time.Time      `json:"subscription_info_updated_at"`
	ProfilePhotoLink             interface{}    `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}    `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}    `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}    `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetMoniestByMoniestId(ctx context.Context, id string) (GetMoniestByMoniestIdRow, error) {
	row := q.db.QueryRowContext(ctx, getMoniestByMoniestId, id)
	var i GetMoniestByMoniestIdRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.SubscriptionInfoID,
		&i.Fee,
		&i.Message,
		&i.SubscriptionInfoUpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const getMoniestByUserId = `-- name: GetMoniestByUserId :one
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
    AND "user"."deleted" = false
`

type GetMoniestByUserIdRow struct {
	ID                           string         `json:"id"`
	MoniestID                    string         `json:"moniest_id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	Bio                          sql.NullString `json:"bio"`
	Description                  sql.NullString `json:"description"`
	Score                        float64        `json:"score"`
	SubscriptionInfoID           string         `json:"subscription_info_id"`
	Fee                          float64        `json:"fee"`
	Message                      sql.NullString `json:"message"`
	SubscriptionInfoUpdatedAt    time.Time      `json:"subscription_info_updated_at"`
	ProfilePhotoLink             interface{}    `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}    `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}    `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}    `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetMoniestByUserId(ctx context.Context, userID string) (GetMoniestByUserIdRow, error) {
	row := q.db.QueryRowContext(ctx, getMoniestByUserId, userID)
	var i GetMoniestByUserIdRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.SubscriptionInfoID,
		&i.Fee,
		&i.Message,
		&i.SubscriptionInfoUpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const updateMoniest = `-- name: UpdateMoniest :one
UPDATE moniest
SET bio = $2,
    description = $3,
    updated_at = now()
WHERE moniest.id = $1
RETURNING id, user_id, bio, description, score, deleted, created_at, updated_at
`

type UpdateMoniestParams struct {
	ID          string         `json:"id"`
	Bio         sql.NullString `json:"bio"`
	Description sql.NullString `json:"description"`
}

// -- name: DeleteMoniest :one
// UPDATE moniest
// SET deleted = true,
//
//	updated_at = now()
//
// WHERE moniest.id = $1
// RETURNING *;
func (q *Queries) UpdateMoniest(ctx context.Context, arg UpdateMoniestParams) (Moniest, error) {
	row := q.db.QueryRowContext(ctx, updateMoniest, arg.ID, arg.Bio, arg.Description)
	var i Moniest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
