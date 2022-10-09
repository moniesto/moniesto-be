// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: moniest.sql

package db

import (
	"context"
	"database/sql"
)

const createMoniest = `-- name: CreateMoniest :one
INSERT INTO "moniest" (
        id,
        user_id,
        bio,
        description
    )
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, bio, description, score, created_at, updated_at
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
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
