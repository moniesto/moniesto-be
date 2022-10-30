// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: login.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const loginUserByEmail = `-- name: LoginUserByEmail :one
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
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE "user"."email" = $1
`

type LoginUserByEmailRow struct {
	ID                           string          `json:"id"`
	MoniestID                    sql.NullString  `json:"moniest_id"`
	Name                         string          `json:"name"`
	Surname                      string          `json:"surname"`
	Username                     string          `json:"username"`
	Email                        string          `json:"email"`
	EmailVerified                bool            `json:"email_verified"`
	Password                     string          `json:"password"`
	Location                     sql.NullString  `json:"location"`
	CreatedAt                    time.Time       `json:"created_at"`
	UpdatedAt                    time.Time       `json:"updated_at"`
	Bio                          sql.NullString  `json:"bio"`
	Description                  sql.NullString  `json:"description"`
	Score                        sql.NullFloat64 `json:"score"`
	ProfilePhotoLink             interface{}     `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}     `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}     `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}     `json:"background_photo_thumbnail_link"`
}

func (q *Queries) LoginUserByEmail(ctx context.Context, email string) (LoginUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, loginUserByEmail, email)
	var i LoginUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const loginUserByUsername = `-- name: LoginUserByUsername :one
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
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
WHERE "user"."username" = $1
`

type LoginUserByUsernameRow struct {
	ID                           string          `json:"id"`
	MoniestID                    sql.NullString  `json:"moniest_id"`
	Name                         string          `json:"name"`
	Surname                      string          `json:"surname"`
	Username                     string          `json:"username"`
	Email                        string          `json:"email"`
	EmailVerified                bool            `json:"email_verified"`
	Password                     string          `json:"password"`
	Location                     sql.NullString  `json:"location"`
	CreatedAt                    time.Time       `json:"created_at"`
	UpdatedAt                    time.Time       `json:"updated_at"`
	Bio                          sql.NullString  `json:"bio"`
	Description                  sql.NullString  `json:"description"`
	Score                        sql.NullFloat64 `json:"score"`
	ProfilePhotoLink             interface{}     `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    interface{}     `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          interface{}     `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink interface{}     `json:"background_photo_thumbnail_link"`
}

func (q *Queries) LoginUserByUsername(ctx context.Context, username string) (LoginUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, loginUserByUsername, username)
	var i LoginUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Bio,
		&i.Description,
		&i.Score,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}
