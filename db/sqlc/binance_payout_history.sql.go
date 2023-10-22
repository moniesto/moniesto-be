// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: binance_payout_history.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createBinancePayoutHistory = `-- name: CreateBinancePayoutHistory :one
INSERT INTO "binance_payout_history" (
        id,
        transaction_id,
        user_id,
        moniest_id,
        payer_id,
        total_amount,
        amount,
        date_type,
        date_value,
        date_index,
        payout_date,
        payout_year,
        payout_month,
        payout_day,
        status,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        now(),
        now()
    )
RETURNING id, transaction_id, user_id, moniest_id, payer_id, total_amount, amount, date_type, date_value, date_index, payout_date, payout_year, payout_month, payout_day, status, operation_fee_percentage, payout_done_at, payout_request_id, failure_message, created_at, updated_at, request, response
`

type CreateBinancePayoutHistoryParams struct {
	ID            string                 `json:"id"`
	TransactionID string                 `json:"transaction_id"`
	UserID        string                 `json:"user_id"`
	MoniestID     string                 `json:"moniest_id"`
	PayerID       string                 `json:"payer_id"`
	TotalAmount   float64                `json:"total_amount"`
	Amount        float64                `json:"amount"`
	DateType      BinancePaymentDateType `json:"date_type"`
	DateValue     int32                  `json:"date_value"`
	DateIndex     int32                  `json:"date_index"`
	PayoutDate    time.Time              `json:"payout_date"`
	PayoutYear    int32                  `json:"payout_year"`
	PayoutMonth   int32                  `json:"payout_month"`
	PayoutDay     int32                  `json:"payout_day"`
	Status        BinancePayoutStatus    `json:"status"`
}

func (q *Queries) CreateBinancePayoutHistory(ctx context.Context, arg CreateBinancePayoutHistoryParams) (BinancePayoutHistory, error) {
	row := q.db.QueryRowContext(ctx, createBinancePayoutHistory,
		arg.ID,
		arg.TransactionID,
		arg.UserID,
		arg.MoniestID,
		arg.PayerID,
		arg.TotalAmount,
		arg.Amount,
		arg.DateType,
		arg.DateValue,
		arg.DateIndex,
		arg.PayoutDate,
		arg.PayoutYear,
		arg.PayoutMonth,
		arg.PayoutDay,
		arg.Status,
	)
	var i BinancePayoutHistory
	err := row.Scan(
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
		&i.PayoutDoneAt,
		&i.PayoutRequestID,
		&i.FailureMessage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Request,
		&i.Response,
	)
	return i, err
}

const getBinancePayoutHistories = `-- name: GetBinancePayoutHistories :many
SELECT "id",
    "transaction_id",
    "user_id",
    "moniest_id",
    "payer_id",
    "total_amount",
    "amount",
    "date_type",
    "date_value",
    "date_index",
    "payout_date",
    "payout_year",
    "payout_month",
    "payout_day",
    "status",
    "operation_fee_percentage",
    "payout_done_at",
    "payout_request_id",
    "failure_message"
FROM "binance_payout_history"
WHERE "transaction_id" = $1
    AND "user_id" = $2
    AND "moniest_id" = $3
    AND "status" = 'pending'
ORDER BY date_index ASC
`

type GetBinancePayoutHistoriesParams struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	MoniestID     string `json:"moniest_id"`
}

type GetBinancePayoutHistoriesRow struct {
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
	PayoutDoneAt           sql.NullTime           `json:"payout_done_at"`
	PayoutRequestID        sql.NullString         `json:"payout_request_id"`
	FailureMessage         sql.NullString         `json:"failure_message"`
}

func (q *Queries) GetBinancePayoutHistories(ctx context.Context, arg GetBinancePayoutHistoriesParams) ([]GetBinancePayoutHistoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getBinancePayoutHistories, arg.TransactionID, arg.UserID, arg.MoniestID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetBinancePayoutHistoriesRow{}
	for rows.Next() {
		var i GetBinancePayoutHistoriesRow
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
			&i.PayoutDoneAt,
			&i.PayoutRequestID,
			&i.FailureMessage,
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

const updateBinancePayoutHistoryRefund = `-- name: UpdateBinancePayoutHistoryRefund :exec
UPDATE "binance_payout_history"
SET "status" = $2,
    "failure_message" = $3,
    "request" = $4,
    "response" = $5,
    updated_at = now()
WHERE id = $1
`

type UpdateBinancePayoutHistoryRefundParams struct {
	ID             string              `json:"id"`
	Status         BinancePayoutStatus `json:"status"`
	FailureMessage sql.NullString      `json:"failure_message"`
	Request        sql.NullString      `json:"request"`
	Response       sql.NullString      `json:"response"`
}

func (q *Queries) UpdateBinancePayoutHistoryRefund(ctx context.Context, arg UpdateBinancePayoutHistoryRefundParams) error {
	_, err := q.db.ExecContext(ctx, updateBinancePayoutHistoryRefund,
		arg.ID,
		arg.Status,
		arg.FailureMessage,
		arg.Request,
		arg.Response,
	)
	return err
}
