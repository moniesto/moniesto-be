// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: content.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getDeactivePosts = `-- name: GetDeactivePosts :many
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
LIMIT $1 OFFSET $2
`

type GetDeactivePostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetDeactivePostsRow struct {
	ID                           string           `json:"id"`
	Currency                     string           `json:"currency"`
	StartPrice                   float64          `json:"start_price"`
	Duration                     time.Time        `json:"duration"`
	Target1                      float64          `json:"target1"`
	Target2                      float64          `json:"target2"`
	Target3                      float64          `json:"target3"`
	Stop                         float64          `json:"stop"`
	Direction                    EntryPosition    `json:"direction"`
	Finished                     bool             `json:"finished"`
	Status                       PostCryptoStatus `json:"status"`
	CreatedAt                    time.Time        `json:"created_at"`
	UpdatedAt                    time.Time        `json:"updated_at"`
	MoniestID                    string           `json:"moniest_id"`
	Bio                          sql.NullString   `json:"bio"`
	Description                  sql.NullString   `json:"description"`
	MoniestScore                 float64          `json:"moniest_score"`
	UserID                       string           `json:"user_id"`
	Name                         string           `json:"name"`
	Surname                      string           `json:"surname"`
	Username                     string           `json:"username"`
	EmailVerified                bool             `json:"email_verified"`
	Location                     sql.NullString   `json:"location"`
	ProfilePhotoLink             interface{}      `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}      `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}      `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}      `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetDeactivePosts(ctx context.Context, arg GetDeactivePostsParams) ([]GetDeactivePostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDeactivePosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetDeactivePostsRow{}
	for rows.Next() {
		var i GetDeactivePostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Currency,
			&i.StartPrice,
			&i.Duration,
			&i.Target1,
			&i.Target2,
			&i.Target3,
			&i.Stop,
			&i.Direction,
			&i.Finished,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MoniestID,
			&i.Bio,
			&i.Description,
			&i.MoniestScore,
			&i.UserID,
			&i.Name,
			&i.Surname,
			&i.Username,
			&i.EmailVerified,
			&i.Location,
			&i.ProfilePhotoLink,
			&i.ProfilePhotoThumbnailLink,
			&i.BackgroundPhotoLink,
			&i.BackgroundPhotoThumbnailLink,
		); err != nil {
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

const getSubscribedActivePosts = `-- name: GetSubscribedActivePosts :many
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
LIMIT $2 OFFSET $3
`

type GetSubscribedActivePostsParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetSubscribedActivePostsRow struct {
	ID                           string           `json:"id"`
	Currency                     string           `json:"currency"`
	StartPrice                   float64          `json:"start_price"`
	Duration                     time.Time        `json:"duration"`
	Target1                      float64          `json:"target1"`
	Target2                      float64          `json:"target2"`
	Target3                      float64          `json:"target3"`
	Stop                         float64          `json:"stop"`
	Direction                    EntryPosition    `json:"direction"`
	Finished                     bool             `json:"finished"`
	Status                       PostCryptoStatus `json:"status"`
	CreatedAt                    time.Time        `json:"created_at"`
	UpdatedAt                    time.Time        `json:"updated_at"`
	MoniestID                    string           `json:"moniest_id"`
	Bio                          sql.NullString   `json:"bio"`
	Description                  sql.NullString   `json:"description"`
	MoniestScore                 float64          `json:"moniest_score"`
	UserID                       string           `json:"user_id"`
	Name                         string           `json:"name"`
	Surname                      string           `json:"surname"`
	Username                     string           `json:"username"`
	EmailVerified                bool             `json:"email_verified"`
	Location                     sql.NullString   `json:"location"`
	ProfilePhotoLink             interface{}      `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}      `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}      `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}      `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetSubscribedActivePosts(ctx context.Context, arg GetSubscribedActivePostsParams) ([]GetSubscribedActivePostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSubscribedActivePosts, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSubscribedActivePostsRow{}
	for rows.Next() {
		var i GetSubscribedActivePostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Currency,
			&i.StartPrice,
			&i.Duration,
			&i.Target1,
			&i.Target2,
			&i.Target3,
			&i.Stop,
			&i.Direction,
			&i.Finished,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MoniestID,
			&i.Bio,
			&i.Description,
			&i.MoniestScore,
			&i.UserID,
			&i.Name,
			&i.Surname,
			&i.Username,
			&i.EmailVerified,
			&i.Location,
			&i.ProfilePhotoLink,
			&i.ProfilePhotoThumbnailLink,
			&i.BackgroundPhotoLink,
			&i.BackgroundPhotoThumbnailLink,
		); err != nil {
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

const getSubscribedDeactivePosts = `-- name: GetSubscribedDeactivePosts :many
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
LIMIT $2 OFFSET $3
`

type GetSubscribedDeactivePostsParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetSubscribedDeactivePostsRow struct {
	ID                           string           `json:"id"`
	Currency                     string           `json:"currency"`
	StartPrice                   float64          `json:"start_price"`
	Duration                     time.Time        `json:"duration"`
	Target1                      float64          `json:"target1"`
	Target2                      float64          `json:"target2"`
	Target3                      float64          `json:"target3"`
	Stop                         float64          `json:"stop"`
	Direction                    EntryPosition    `json:"direction"`
	Finished                     bool             `json:"finished"`
	Status                       PostCryptoStatus `json:"status"`
	CreatedAt                    time.Time        `json:"created_at"`
	UpdatedAt                    time.Time        `json:"updated_at"`
	MoniestID                    string           `json:"moniest_id"`
	Bio                          sql.NullString   `json:"bio"`
	Description                  sql.NullString   `json:"description"`
	MoniestScore                 float64          `json:"moniest_score"`
	UserID                       string           `json:"user_id"`
	Name                         string           `json:"name"`
	Surname                      string           `json:"surname"`
	Username                     string           `json:"username"`
	EmailVerified                bool             `json:"email_verified"`
	Location                     sql.NullString   `json:"location"`
	ProfilePhotoLink             interface{}      `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}      `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}      `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}      `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetSubscribedDeactivePosts(ctx context.Context, arg GetSubscribedDeactivePostsParams) ([]GetSubscribedDeactivePostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSubscribedDeactivePosts, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSubscribedDeactivePostsRow{}
	for rows.Next() {
		var i GetSubscribedDeactivePostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Currency,
			&i.StartPrice,
			&i.Duration,
			&i.Target1,
			&i.Target2,
			&i.Target3,
			&i.Stop,
			&i.Direction,
			&i.Finished,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MoniestID,
			&i.Bio,
			&i.Description,
			&i.MoniestScore,
			&i.UserID,
			&i.Name,
			&i.Surname,
			&i.Username,
			&i.EmailVerified,
			&i.Location,
			&i.ProfilePhotoLink,
			&i.ProfilePhotoThumbnailLink,
			&i.BackgroundPhotoLink,
			&i.BackgroundPhotoThumbnailLink,
		); err != nil {
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

const getSubscribedMoniests = `-- name: GetSubscribedMoniests :many
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
LIMIT $2 OFFSET $3
`

type GetSubscribedMoniestsParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetSubscribedMoniestsRow struct {
	ID                           string         `json:"id"`
	MoniestID                    string         `json:"moniest_id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	EmailVerified                bool           `json:"email_verified"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	Bio                          sql.NullString `json:"bio"`
	Description                  sql.NullString `json:"description"`
	Score                        float64        `json:"score"`
	Fee                          float64        `json:"fee"`
	Message                      sql.NullString `json:"message"`
	SubscriptionInfoUpdatedAt    time.Time      `json:"subscription_info_updated_at"`
	ProfilePhotoLink             interface{}    `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}    `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}    `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}    `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetSubscribedMoniests(ctx context.Context, arg GetSubscribedMoniestsParams) ([]GetSubscribedMoniestsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSubscribedMoniests, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSubscribedMoniestsRow{}
	for rows.Next() {
		var i GetSubscribedMoniestsRow
		if err := rows.Scan(
			&i.ID,
			&i.MoniestID,
			&i.Name,
			&i.Surname,
			&i.Username,
			&i.EmailVerified,
			&i.Location,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Bio,
			&i.Description,
			&i.Score,
			&i.Fee,
			&i.Message,
			&i.SubscriptionInfoUpdatedAt,
			&i.ProfilePhotoLink,
			&i.ProfilePhotoThumbnailLink,
			&i.BackgroundPhotoLink,
			&i.BackgroundPhotoThumbnailLink,
		); err != nil {
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
