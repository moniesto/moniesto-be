// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: moniest_payout_info.sql

package db

import (
	"context"
)

const createMoniestPayoutInfo = `-- name: CreateMoniestPayoutInfo :one
INSERT INTO moniest_payout_info (
        id,
        moniest_id,
        source,
        type,
        value,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        now(),
        now()
    )
RETURNING id, moniest_id, source, type, value, created_at, updated_at
`

type CreateMoniestPayoutInfoParams struct {
	ID        string       `json:"id"`
	MoniestID string       `json:"moniest_id"`
	Source    PayoutSource `json:"source"`
	Type      PayoutType   `json:"type"`
	Value     string       `json:"value"`
}

func (q *Queries) CreateMoniestPayoutInfo(ctx context.Context, arg CreateMoniestPayoutInfoParams) (MoniestPayoutInfo, error) {
	row := q.db.QueryRowContext(ctx, createMoniestPayoutInfo,
		arg.ID,
		arg.MoniestID,
		arg.Source,
		arg.Type,
		arg.Value,
	)
	var i MoniestPayoutInfo
	err := row.Scan(
		&i.ID,
		&i.MoniestID,
		&i.Source,
		&i.Type,
		&i.Value,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
