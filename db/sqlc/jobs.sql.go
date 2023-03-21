// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: jobs.sql

package db

import (
	"context"
	"time"
)

const getAllActivePosts = `-- name: GetAllActivePosts :many
SELECT "pc"."id",
    "pc"."moniest_id",
    "pc"."currency",
    "pc"."start_price",
    "pc"."duration",
    "pc"."target1",
    "pc"."target2",
    "pc"."target3",
    "pc"."stop",
    "pc"."direction",
    "pc"."score",
    "pc"."finished",
    "pc"."status",
    "pc"."last_target_hit",
    "pc"."last_job_timestamp",
    "pc"."created_at",
    "pc"."updated_at"
FROM "post_crypto" AS pc
WHERE "pc"."duration" > now()
    AND "pc"."finished" = FALSE
ORDER BY "pc"."created_at" ASC
`

type GetAllActivePostsRow struct {
	ID               string           `json:"id"`
	MoniestID        string           `json:"moniest_id"`
	Currency         string           `json:"currency"`
	StartPrice       float64          `json:"start_price"`
	Duration         time.Time        `json:"duration"`
	Target1          float64          `json:"target1"`
	Target2          float64          `json:"target2"`
	Target3          float64          `json:"target3"`
	Stop             float64          `json:"stop"`
	Direction        EntryPosition    `json:"direction"`
	Score            float64          `json:"score"`
	Finished         bool             `json:"finished"`
	Status           PostCryptoStatus `json:"status"`
	LastTargetHit    float64          `json:"last_target_hit"`
	LastJobTimestamp int64            `json:"last_job_timestamp"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

func (q *Queries) GetAllActivePosts(ctx context.Context) ([]GetAllActivePostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllActivePosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllActivePostsRow{}
	for rows.Next() {
		var i GetAllActivePostsRow
		if err := rows.Scan(
			&i.ID,
			&i.MoniestID,
			&i.Currency,
			&i.StartPrice,
			&i.Duration,
			&i.Target1,
			&i.Target2,
			&i.Target3,
			&i.Stop,
			&i.Direction,
			&i.Score,
			&i.Finished,
			&i.Status,
			&i.LastTargetHit,
			&i.LastJobTimestamp,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateFinishedPostStatus = `-- name: UpdateFinishedPostStatus :exec
UPDATE "post_crypto"
SET "status" = $1,
    "score" = $2,
    "finished" = TRUE
WHERE "id" = $3
`

type UpdateFinishedPostStatusParams struct {
	Status PostCryptoStatus `json:"status"`
	Score  float64          `json:"score"`
	ID     string           `json:"id"`
}

func (q *Queries) UpdateFinishedPostStatus(ctx context.Context, arg UpdateFinishedPostStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateFinishedPostStatus, arg.Status, arg.Score, arg.ID)
	return err
}

const updateMoniestScore = `-- name: UpdateMoniestScore :exec
UPDATE "moniest"
SET "score" = GREATEST("score" + $1, 0)
WHERE "id" = $2
`

type UpdateMoniestScoreParams struct {
	Score float64 `json:"score"`
	ID    string  `json:"id"`
}

func (q *Queries) UpdateMoniestScore(ctx context.Context, arg UpdateMoniestScoreParams) error {
	_, err := q.db.ExecContext(ctx, updateMoniestScore, arg.Score, arg.ID)
	return err
}

const updateUnfinishedPostStatus = `-- name: UpdateUnfinishedPostStatus :exec
UPDATE "post_crypto"
SET "last_target_hit" = $1,
    "last_job_timestamp" = $2
WHERE "id" = $3
`

type UpdateUnfinishedPostStatusParams struct {
	LastTargetHit    float64 `json:"last_target_hit"`
	LastJobTimestamp int64   `json:"last_job_timestamp"`
	ID               string  `json:"id"`
}

func (q *Queries) UpdateUnfinishedPostStatus(ctx context.Context, arg UpdateUnfinishedPostStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateUnfinishedPostStatus, arg.LastTargetHit, arg.LastJobTimestamp, arg.ID)
	return err
}
