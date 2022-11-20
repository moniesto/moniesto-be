// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: password_reset_token.sql

package db

import (
	"context"
	"time"
)

const createPasswordResetToken = `-- name: CreatePasswordResetToken :one
INSERT INTO "password_reset_token" (
        id,
        user_id,
        token,
        token_expiry,
        deleted,
        created_at
    )
VALUEs ($1, $2, $3, $4, false, now())
RETURNING id, user_id, token, token_expiry, deleted, created_at
`

type CreatePasswordResetTokenParams struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
}

func (q *Queries) CreatePasswordResetToken(ctx context.Context, arg CreatePasswordResetTokenParams) (PasswordResetToken, error) {
	row := q.db.QueryRowContext(ctx, createPasswordResetToken,
		arg.ID,
		arg.UserID,
		arg.Token,
		arg.TokenExpiry,
	)
	var i PasswordResetToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.TokenExpiry,
		&i.Deleted,
		&i.CreatedAt,
	)
	return i, err
}

const deletePasswordResetTokenByToken = `-- name: DeletePasswordResetTokenByToken :exec
UPDATE "password_reset_token"
SET deleted = true
WHERE "token" = $1
`

func (q *Queries) DeletePasswordResetTokenByToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deletePasswordResetTokenByToken, token)
	return err
}

const getPasswordResetTokenByToken = `-- name: GetPasswordResetTokenByToken :one
SELECT id, user_id, token, token_expiry, deleted, created_at
FROM "password_reset_token"
WHERE "token" = $1
    AND "deleted" = false
`

func (q *Queries) GetPasswordResetTokenByToken(ctx context.Context, token string) (PasswordResetToken, error) {
	row := q.db.QueryRowContext(ctx, getPasswordResetTokenByToken, token)
	var i PasswordResetToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.TokenExpiry,
		&i.Deleted,
		&i.CreatedAt,
	)
	return i, err
}
