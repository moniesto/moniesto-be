// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const checkEmail = `-- name: CheckEmail :one
SELECT COUNT(*) = 0 AS isEmailValid
FROM "user"
WHERE email = $1
`

func (q *Queries) CheckEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkEmail, email)
	var isemailvalid bool
	err := row.Scan(&isemailvalid)
	return isemailvalid, err
}

const checkUsername = `-- name: CheckUsername :one
SELECT COUNT(*) = 0 AS isUsernameValid
FROM "user"
WHERE username = $1
`

func (q *Queries) CheckUsername(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUsername, username)
	var isusernamevalid bool
	err := row.Scan(&isusernamevalid)
	return isusernamevalid, err
}

const createUser = `-- name: CreateUser :one
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
RETURNING id, name, surname, username, email, email_verified, password, location, login_count, deleted, created_at, updated_at, last_login
`

type CreateUserParams struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Surname,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.LoginCount,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
UPDATE "user"
SET deleted = true,
    updated_at = now()
WHERE id = $1
RETURNING id, name, surname, username, email, email_verified, password, location, login_count, deleted, created_at, updated_at, last_login
`

func (q *Queries) DeleteUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.LoginCount,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}

const getActiveUsersVerifiedEmails = `-- name: GetActiveUsersVerifiedEmails :many
SELECT email
FROM "user"
WHERE email_verified = true
    AND deleted = false
`

func (q *Queries) GetActiveUsersVerifiedEmails(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getActiveUsersVerifiedEmails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInactiveUsersVerifiedEmails = `-- name: GetInactiveUsersVerifiedEmails :many
SELECT email
FROM "user"
WHERE email_verified = true
    AND deleted = true
`

func (q *Queries) GetInactiveUsersVerifiedEmails(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getInactiveUsersVerifiedEmails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT "user"."id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_thumbnail_link",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_thumbnail_link"
FROM "user"
WHERE "user"."email" = $1
`

type GetUserByEmailRow struct {
	ID                           string         `json:"id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	ProfilePhotoLink             string         `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    string         `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          string         `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink string         `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT "user"."id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    (
        SELECT "image"."link"
        FROM "image"
        WHERE "image"."user_id" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
        WHERE "image"."user_id" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_thumbnail_link",
    (
        SELECT "image"."link"
        FROM "image"
        WHERE "image"."user_id" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
        WHERE "image"."user_id" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_thumbnail_link"
FROM "user"
WHERE "user"."id" = $1
`

type GetUserByIDRow struct {
	ID                           string         `json:"id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	ProfilePhotoLink             string         `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    string         `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          string         `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink string         `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetUserByID(ctx context.Context, userID string) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT "user"."id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_thumbnail_link",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_thumbnail_link"
FROM "user"
WHERE "user"."username" = $1
`

type GetUserByUsernameRow struct {
	ID                           string         `json:"id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	ProfilePhotoLink             string         `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    string         `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          string         `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink string         `json:"background_photo_thumbnail_link"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const loginUserByEmail = `-- name: LoginUserByEmail :one
SELECT "user"."id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."password",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_thumbnail_link",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."email" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_thumbnail_link"
FROM "user"
WHERE "user"."email" = $1
`

type LoginUserByEmailRow struct {
	ID                           string         `json:"id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Password                     string         `json:"password"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	ProfilePhotoLink             string         `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    string         `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          string         `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink string         `json:"background_photo_thumbnail_link"`
}

func (q *Queries) LoginUserByEmail(ctx context.Context, email string) (LoginUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, loginUserByEmail, email)
	var i LoginUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const loginUserByUsername = `-- name: LoginUserByUsername :one
SELECT "user"."id",
    "user"."name",
    "user"."surname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."password",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "profile_photo"
    ) AS "profile_photo_thumbnail_link",
    (
        SELECT "image"."link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_link",
    (
        SELECT "image"."thumbnail_link"
        FROM "image"
            INNER JOIN "user" ON "user"."id" = "image"."user_id"
        WHERE "user"."username" = $1
            AND "image"."type" = "background_photo"
    ) AS "background_photo_thumbnail_link"
FROM "user"
WHERE "user"."username" = $1
`

type LoginUserByUsernameRow struct {
	ID                           string         `json:"id"`
	Name                         string         `json:"name"`
	Surname                      string         `json:"surname"`
	Username                     string         `json:"username"`
	Email                        string         `json:"email"`
	EmailVerified                bool           `json:"email_verified"`
	Password                     string         `json:"password"`
	Location                     sql.NullString `json:"location"`
	CreatedAt                    time.Time      `json:"created_at"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	ProfilePhotoLink             string         `json:"profile_photo_link"`
	ProfilePhotoThumbnailLink    string         `json:"profile_photo_thumbnail_link"`
	BackgroundPhotoLink          string         `json:"background_photo_link"`
	BackgroundPhotoThumbnailLink string         `json:"background_photo_thumbnail_link"`
}

func (q *Queries) LoginUserByUsername(ctx context.Context, username string) (LoginUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, loginUserByUsername, username)
	var i LoginUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProfilePhotoLink,
		&i.ProfilePhotoThumbnailLink,
		&i.BackgroundPhotoLink,
		&i.BackgroundPhotoThumbnailLink,
	)
	return i, err
}

const setPassword = `-- name: SetPassword :one
UPDATE "user"
SET password = $2
WHERE id = $1
RETURNING id, name, surname, username, email, email_verified, password, location, login_count, deleted, created_at, updated_at, last_login
`

type SetPasswordParams struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) SetPassword(ctx context.Context, arg SetPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, setPassword, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.LoginCount,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}

const updateLoginStats = `-- name: UpdateLoginStats :one
UPDATE "user"
SET login_count = login_count + 1,
    last_login = now()
WHERE id = $1
RETURNING id, name, surname, username, email, email_verified, password, location, login_count, deleted, created_at, updated_at, last_login
`

func (q *Queries) UpdateLoginStats(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, updateLoginStats, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Username,
		&i.Email,
		&i.EmailVerified,
		&i.Password,
		&i.Location,
		&i.LoginCount,
		&i.Deleted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}
