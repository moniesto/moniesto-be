// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: jobs.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createUserSubscriptionHistory = `-- name: CreateUserSubscriptionHistory :one
INSERT INTO "user_subscription_history" (
        id,
        user_id,
        moniest_id,
        transaction_id,
        subscription_start_date,
        subscription_end_date,
        created_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        now()
    )
RETURNING id, user_id, moniest_id, transaction_id, subscription_start_date, subscription_end_date, created_at
`

type CreateUserSubscriptionHistoryParams struct {
	ID                    string         `json:"id"`
	UserID                string         `json:"user_id"`
	MoniestID             string         `json:"moniest_id"`
	TransactionID         sql.NullString `json:"transaction_id"`
	SubscriptionStartDate time.Time      `json:"subscription_start_date"`
	SubscriptionEndDate   time.Time      `json:"subscription_end_date"`
}

func (q *Queries) CreateUserSubscriptionHistory(ctx context.Context, arg CreateUserSubscriptionHistoryParams) (UserSubscriptionHistory, error) {
	row := q.db.QueryRowContext(ctx, createUserSubscriptionHistory,
		arg.ID,
		arg.UserID,
		arg.MoniestID,
		arg.TransactionID,
		arg.SubscriptionStartDate,
		arg.SubscriptionEndDate,
	)
	var i UserSubscriptionHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MoniestID,
		&i.TransactionID,
		&i.SubscriptionStartDate,
		&i.SubscriptionEndDate,
		&i.CreatedAt,
	)
	return i, err
}

const getAllActivePosts = `-- name: GetAllActivePosts :many
SELECT "pc"."id",
    "pc"."moniest_id",
    "pc"."market_type",
    "pc"."currency",
    "pc"."start_price",
    "pc"."duration",
    "pc"."take_profit",
    "pc"."stop",
    "pc"."target1",
    "pc"."target2",
    "pc"."target3",
    "pc"."direction",
    "pc"."leverage",
    "pc"."finished",
    "pc"."status",
    "pc"."pnl",
    "pc"."roi",
    "pc"."last_operated_at",
    "pc"."created_at",
    "pc"."updated_at"
FROM "post_crypto" AS pc
WHERE "pc"."finished" = FALSE
ORDER BY "pc"."created_at" ASC
`

type GetAllActivePostsRow struct {
	ID             string               `json:"id"`
	MoniestID      string               `json:"moniest_id"`
	MarketType     PostCryptoMarketType `json:"market_type"`
	Currency       string               `json:"currency"`
	StartPrice     float64              `json:"start_price"`
	Duration       time.Time            `json:"duration"`
	TakeProfit     float64              `json:"take_profit"`
	Stop           float64              `json:"stop"`
	Target1        sql.NullFloat64      `json:"target1"`
	Target2        sql.NullFloat64      `json:"target2"`
	Target3        sql.NullFloat64      `json:"target3"`
	Direction      EntryPosition        `json:"direction"`
	Leverage       int32                `json:"leverage"`
	Finished       bool                 `json:"finished"`
	Status         PostCryptoStatus     `json:"status"`
	Pnl            float64              `json:"pnl"`
	Roi            float64              `json:"roi"`
	LastOperatedAt time.Time            `json:"last_operated_at"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
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
			&i.MarketType,
			&i.Currency,
			&i.StartPrice,
			&i.Duration,
			&i.TakeProfit,
			&i.Stop,
			&i.Target1,
			&i.Target2,
			&i.Target3,
			&i.Direction,
			&i.Leverage,
			&i.Finished,
			&i.Status,
			&i.Pnl,
			&i.Roi,
			&i.LastOperatedAt,
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

const getAllPendingPayouts = `-- name: GetAllPendingPayouts :many
SELECT "bph"."id",
    "bph"."transaction_id",
    "bph"."user_id",
    "bph"."moniest_id",
    "bph"."payer_id",
    "bph"."total_amount",
    "bph"."amount",
    "bph"."date_type",
    "bph"."date_value",
    "bph"."date_index",
    "bph"."payout_date",
    "bph"."payout_year",
    "bph"."payout_month",
    "bph"."payout_day",
    "bph"."status",
    "bph"."operation_fee_percentage",
    "bph"."created_at",
    "bph"."updated_at",
    "mpi"."type" as "moniest_payout_type",
    "mpi"."value" as "moniest_payout_value"
FROM "binance_payout_history" as "bph"
    INNER JOIN "moniest_payout_info" as "mpi" ON "mpi"."moniest_id" = "bph"."moniest_id"
WHERE "status" = 'pending'
    AND payout_date <= now()
`

type GetAllPendingPayoutsRow struct {
	ID                     string                 `json:"id"`
	TransactionID          string                 `json:"transaction_id"`
	UserID                 string                 `json:"user_id"`
	MoniestID              string                 `json:"moniest_id"`
	PayerID                string                 `json:"payer_id"`
	TotalAmount            float64                `json:"total_amount"`
	Amount                 float64                `json:"amount"`
	DateType               BinancePaymentDateType `json:"date_type"`
	DateValue              int32                  `json:"date_value"`
	DateIndex              int32                  `json:"date_index"`
	PayoutDate             time.Time              `json:"payout_date"`
	PayoutYear             int32                  `json:"payout_year"`
	PayoutMonth            int32                  `json:"payout_month"`
	PayoutDay              int32                  `json:"payout_day"`
	Status                 BinancePayoutStatus    `json:"status"`
	OperationFeePercentage sql.NullFloat64        `json:"operation_fee_percentage"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
	MoniestPayoutType      PayoutType             `json:"moniest_payout_type"`
	MoniestPayoutValue     string                 `json:"moniest_payout_value"`
}

func (q *Queries) GetAllPendingPayouts(ctx context.Context) ([]GetAllPendingPayoutsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllPendingPayouts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllPendingPayoutsRow{}
	for rows.Next() {
		var i GetAllPendingPayoutsRow
		if err := rows.Scan(
			&i.ID,
			&i.TransactionID,
			&i.UserID,
			&i.MoniestID,
			&i.PayerID,
			&i.TotalAmount,
			&i.Amount,
			&i.DateType,
			&i.DateValue,
			&i.DateIndex,
			&i.PayoutDate,
			&i.PayoutYear,
			&i.PayoutMonth,
			&i.PayoutDay,
			&i.Status,
			&i.OperationFeePercentage,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MoniestPayoutType,
			&i.MoniestPayoutValue,
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

const getExpiredActiveSubscriptions = `-- name: GetExpiredActiveSubscriptions :many
SELECT id, user_id, moniest_id, active, latest_transaction_id, subscription_start_date, subscription_end_date, created_at, updated_at
FROM "user_subscription"
WHERE active = TRUE
    AND subscription_end_date <= now()
`

func (q *Queries) GetExpiredActiveSubscriptions(ctx context.Context) ([]UserSubscription, error) {
	rows, err := q.db.QueryContext(ctx, getExpiredActiveSubscriptions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserSubscription{}
	for rows.Next() {
		var i UserSubscription
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.MoniestID,
			&i.Active,
			&i.LatestTransactionID,
			&i.SubscriptionStartDate,
			&i.SubscriptionEndDate,
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

const getExpiredPendingBinanceTransactions = `-- name: GetExpiredPendingBinanceTransactions :many
SELECT id, qrcode_link, checkout_link, deep_link, universal_link, status, user_id, moniest_id, date_type, date_value, moniest_fee, amount, webhook_url, payer_id, created_at, updated_at
FROM binance_payment_transaction
WHERE status = 'pending'
    AND "created_at" + INTERVAL '5 minutes' <= NOW()
`

func (q *Queries) GetExpiredPendingBinanceTransactions(ctx context.Context) ([]BinancePaymentTransaction, error) {
	rows, err := q.db.QueryContext(ctx, getExpiredPendingBinanceTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BinancePaymentTransaction{}
	for rows.Next() {
		var i BinancePaymentTransaction
		if err := rows.Scan(
			&i.ID,
			&i.QrcodeLink,
			&i.CheckoutLink,
			&i.DeepLink,
			&i.UniversalLink,
			&i.Status,
			&i.UserID,
			&i.MoniestID,
			&i.DateType,
			&i.DateValue,
			&i.MoniestFee,
			&i.Amount,
			&i.WebhookUrl,
			&i.PayerID,
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

const updateBinancePayoutHistoryPayout = `-- name: UpdateBinancePayoutHistoryPayout :exec
UPDATE "binance_payout_history"
SET "status" = $2,
    operation_fee_percentage = $3,
    "payout_done_at" = $4,
    "failure_message" = $5,
    payout_request_id = $6,
    updated_at = now()
WHERE "id" = $1
`

type UpdateBinancePayoutHistoryPayoutParams struct {
	ID                     string              `json:"id"`
	Status                 BinancePayoutStatus `json:"status"`
	OperationFeePercentage sql.NullFloat64     `json:"operation_fee_percentage"`
	PayoutDoneAt           sql.NullTime        `json:"payout_done_at"`
	FailureMessage         sql.NullString      `json:"failure_message"`
	PayoutRequestID        sql.NullString      `json:"payout_request_id"`
}

func (q *Queries) UpdateBinancePayoutHistoryPayout(ctx context.Context, arg UpdateBinancePayoutHistoryPayoutParams) error {
	_, err := q.db.ExecContext(ctx, updateBinancePayoutHistoryPayout,
		arg.ID,
		arg.Status,
		arg.OperationFeePercentage,
		arg.PayoutDoneAt,
		arg.FailureMessage,
		arg.PayoutRequestID,
	)
	return err
}

const updateExpiredActiveSubscription = `-- name: UpdateExpiredActiveSubscription :exec
UPDATE "user_subscription"
SET active = FALSE,
    updated_at = now()
WHERE "id" = $1
`

func (q *Queries) UpdateExpiredActiveSubscription(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updateExpiredActiveSubscription, id)
	return err
}

const updateExpiredPendingBinanceTransaction = `-- name: UpdateExpiredPendingBinanceTransaction :exec
UPDATE "binance_payment_transaction"
SET status = 'fail',
    updated_at = now()
WHERE "id" = $1
`

func (q *Queries) UpdateExpiredPendingBinanceTransaction(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updateExpiredPendingBinanceTransaction, id)
	return err
}

const updateFinishedPostStatus = `-- name: UpdateFinishedPostStatus :exec
UPDATE "post_crypto"
SET "status" = $2,
    "pnl" = $3,
    "roi" = $4,
    "finished" = TRUE,
    updated_at = now()
WHERE "id" = $1
`

type UpdateFinishedPostStatusParams struct {
	ID     string           `json:"id"`
	Status PostCryptoStatus `json:"status"`
	Pnl    float64          `json:"pnl"`
	Roi    float64          `json:"roi"`
}

func (q *Queries) UpdateFinishedPostStatus(ctx context.Context, arg UpdateFinishedPostStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateFinishedPostStatus,
		arg.ID,
		arg.Status,
		arg.Pnl,
		arg.Roi,
	)
	return err
}

const updateMoniestPostCryptoStatistics = `-- name: UpdateMoniestPostCryptoStatistics :exec
UPDATE "moniest_post_crypto_statistics"
SET "pnl_7days" = $2,
    "roi_7days" = $3,
    "win_rate_7days" = $4,
    "posts_7days" = $5,
    "pnl_30days" = $6,
    "roi_30days" = $7,
    "win_rate_30days" = $8,
    "posts_30days" = $9,
    "pnl_total" = $10,
    "roi_total" = $11,
    "win_rate_total" = $12,
    "updated_at" = now()
WHERE "moniest_id" = $1
`

type UpdateMoniestPostCryptoStatisticsParams struct {
	MoniestID     string          `json:"moniest_id"`
	Pnl7days      sql.NullFloat64 `json:"pnl_7days"`
	Roi7days      sql.NullFloat64 `json:"roi_7days"`
	WinRate7days  sql.NullFloat64 `json:"win_rate_7days"`
	Posts7days    []string        `json:"posts_7days"`
	Pnl30days     sql.NullFloat64 `json:"pnl_30days"`
	Roi30days     sql.NullFloat64 `json:"roi_30days"`
	WinRate30days sql.NullFloat64 `json:"win_rate_30days"`
	Posts30days   []string        `json:"posts_30days"`
	PnlTotal      sql.NullFloat64 `json:"pnl_total"`
	RoiTotal      sql.NullFloat64 `json:"roi_total"`
	WinRateTotal  sql.NullFloat64 `json:"win_rate_total"`
}

func (q *Queries) UpdateMoniestPostCryptoStatistics(ctx context.Context, arg UpdateMoniestPostCryptoStatisticsParams) error {
	_, err := q.db.ExecContext(ctx, updateMoniestPostCryptoStatistics,
		arg.MoniestID,
		arg.Pnl7days,
		arg.Roi7days,
		arg.WinRate7days,
		pq.Array(arg.Posts7days),
		arg.Pnl30days,
		arg.Roi30days,
		arg.WinRate30days,
		pq.Array(arg.Posts30days),
		arg.PnlTotal,
		arg.RoiTotal,
		arg.WinRateTotal,
	)
	return err
}

const updateUnfinishedPostStatus = `-- name: UpdateUnfinishedPostStatus :exec
UPDATE "post_crypto"
SET "last_operated_at" = $2,
    updated_at = now()
WHERE "id" = $1
`

type UpdateUnfinishedPostStatusParams struct {
	ID             string    `json:"id"`
	LastOperatedAt time.Time `json:"last_operated_at"`
}

func (q *Queries) UpdateUnfinishedPostStatus(ctx context.Context, arg UpdateUnfinishedPostStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateUnfinishedPostStatus, arg.ID, arg.LastOperatedAt)
	return err
}
