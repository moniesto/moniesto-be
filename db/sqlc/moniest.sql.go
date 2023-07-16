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

const checkUserIsMoniestByID = `-- name: CheckUserIsMoniestByID :one
SELECT COUNT(*) != 0 AS userIsMoniest
FROM "moniest"
WHERE "moniest"."user_id" = $1
`

func (q *Queries) CheckUserIsMoniestByID(ctx context.Context, userID string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUserIsMoniestByID, userID)
	var userismoniest bool
	err := row.Scan(&userismoniest)
	return userismoniest, err
}

const checkUserIsMoniestByUsername = `-- name: CheckUserIsMoniestByUsername :one
SELECT COUNT(*) != 0 AS userIsMoniest
FROM "moniest"
    INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
    AND "user"."username" = $1
`

func (q *Queries) CheckUserIsMoniestByUsername(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUserIsMoniestByUsername, username)
	var userismoniest bool
	err := row.Scan(&userismoniest)
	return userismoniest, err
}

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
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    "moniest_subscription_info"."id" as "moniest_subscription_info_id",
    "moniest_subscription_info"."fee",
    "moniest_subscription_info"."message",
    "moniest_subscription_info"."updated_at" as "moniest_subscription_info_updated_at",
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
    INNER JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
WHERE "moniest"."id" = $1
    AND "user"."deleted" = false
`

type GetMoniestByMoniestIdRow struct {
	ID                               string         `json:"id"`
	MoniestID                        string         `json:"moniest_id"`
	Fullname                         string         `json:"fullname"`
	Username                         string         `json:"username"`
	Email                            string         `json:"email"`
	EmailVerified                    bool           `json:"email_verified"`
	Location                         sql.NullString `json:"location"`
	CreatedAt                        time.Time      `json:"created_at"`
	UpdatedAt                        time.Time      `json:"updated_at"`
	Bio                              sql.NullString `json:"bio"`
	Description                      sql.NullString `json:"description"`
	Score                            float64        `json:"score"`
	MoniestSubscriptionInfoID        string         `json:"moniest_subscription_info_id"`
	Fee                              float64        `json:"fee"`
	Message                          sql.NullString `json:"message"`
	MoniestSubscriptionInfoUpdatedAt time.Time      `json:"moniest_subscription_info_updated_at"`
	ProfilePhotoLink                 interface{}    `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink        interface{}    `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink              interface{}    `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink     interface{}    `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetMoniestByMoniestId(ctx context.Context, id string) (GetMoniestByMoniestIdRow, error) {
	row := q.db.QueryRowContext(ctx, getMoniestByMoniestId, id)
	var i GetMoniestByMoniestIdRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Fullname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.MoniestSubscriptionInfoID,
		&i.Fee,
		&i.Message,
		&i.MoniestSubscriptionInfoUpdatedAt,
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
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    "moniest_subscription_info"."id" as "moniest_subscription_info_id",
    "moniest_subscription_info"."fee",
    "moniest_subscription_info"."message",
    "moniest_subscription_info"."updated_at" as "moniest_subscription_info_updated_at",
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
    INNER JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."id" = $1
    AND "user"."deleted" = false
`

type GetMoniestByUserIdRow struct {
	ID                               string         `json:"id"`
	MoniestID                        string         `json:"moniest_id"`
	Fullname                         string         `json:"fullname"`
	Username                         string         `json:"username"`
	Email                            string         `json:"email"`
	EmailVerified                    bool           `json:"email_verified"`
	Location                         sql.NullString `json:"location"`
	CreatedAt                        time.Time      `json:"created_at"`
	UpdatedAt                        time.Time      `json:"updated_at"`
	Bio                              sql.NullString `json:"bio"`
	Description                      sql.NullString `json:"description"`
	Score                            float64        `json:"score"`
	MoniestSubscriptionInfoID        string         `json:"moniest_subscription_info_id"`
	Fee                              float64        `json:"fee"`
	Message                          sql.NullString `json:"message"`
	MoniestSubscriptionInfoUpdatedAt time.Time      `json:"moniest_subscription_info_updated_at"`
	ProfilePhotoLink                 interface{}    `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink        interface{}    `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink              interface{}    `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink     interface{}    `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetMoniestByUserId(ctx context.Context, userID string) (GetMoniestByUserIdRow, error) {
	row := q.db.QueryRowContext(ctx, getMoniestByUserId, userID)
	var i GetMoniestByUserIdRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Fullname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.MoniestSubscriptionInfoID,
		&i.Fee,
		&i.Message,
		&i.MoniestSubscriptionInfoUpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const getMoniestByUsername = `-- name: GetMoniestByUsername :one
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "moniest"."score",
    "moniest_subscription_info"."id" as "moniest_subscription_info_id",
    "moniest_subscription_info"."fee",
    "moniest_subscription_info"."message",
    "moniest_subscription_info"."updated_at" as "moniest_subscription_info_updated_at",
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
    INNER JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
WHERE "user"."username" = $1
    AND "user"."deleted" = false
`

type GetMoniestByUsernameRow struct {
	ID                               string         `json:"id"`
	MoniestID                        string         `json:"moniest_id"`
	Fullname                         string         `json:"fullname"`
	Username                         string         `json:"username"`
	Email                            string         `json:"email"`
	EmailVerified                    bool           `json:"email_verified"`
	Location                         sql.NullString `json:"location"`
	CreatedAt                        time.Time      `json:"created_at"`
	UpdatedAt                        time.Time      `json:"updated_at"`
	Bio                              sql.NullString `json:"bio"`
	Description                      sql.NullString `json:"description"`
	Score                            float64        `json:"score"`
	MoniestSubscriptionInfoID        string         `json:"moniest_subscription_info_id"`
	Fee                              float64        `json:"fee"`
	Message                          sql.NullString `json:"message"`
	MoniestSubscriptionInfoUpdatedAt time.Time      `json:"moniest_subscription_info_updated_at"`
	ProfilePhotoLink                 interface{}    `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink        interface{}    `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink              interface{}    `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink     interface{}    `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetMoniestByUsername(ctx context.Context, username string) (GetMoniestByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getMoniestByUsername, username)
	var i GetMoniestByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Fullname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.MoniestSubscriptionInfoID,
		&i.Fee,
		&i.Message,
		&i.MoniestSubscriptionInfoUpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const getMoniestStatsByUsername = `-- name: GetMoniestStatsByUsername :one
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
where "user"."username" = $1
`

type GetMoniestStatsByUsernameRow struct {
	SubscriptionCount int64 `json:"subscription_count"`
	SubscriberCount   int64 `json:"subscriber_count"`
	PostCount         int64 `json:"post_count"`
}

func (q *Queries) GetMoniestStatsByUsername(ctx context.Context, username string) (GetMoniestStatsByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getMoniestStatsByUsername, username)
	var i GetMoniestStatsByUsernameRow
	err := row.Scan(&i.SubscriptionCount, &i.SubscriberCount, &i.PostCount)
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
