// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: binance_payment_transaction.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const checkPendingBinancePaymentTransactionByMoniestUsername = `-- name: CheckPendingBinancePaymentTransactionByMoniestUsername :many
SELECT COUNT(*) != 0 as pending,
    "binance_payment_transaction"."qrcode_link",
    "binance_payment_transaction"."checkout_link",
    "binance_payment_transaction"."deep_link",
    "binance_payment_transaction"."universal_link",
    "binance_payment_transaction"."created_at"
FROM "binance_payment_transaction"
    INNER JOIN "moniest" ON "moniest"."id" = "binance_payment_transaction"."moniest_id"
    INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
    AND "user"."username" = $2
WHERE "binance_payment_transaction"."status" = 'pending'
    AND "binance_payment_transaction"."user_id" = $1
GROUP BY "binance_payment_transaction"."id"
ORDER BY "binance_payment_transaction"."created_at" DESC
`

type CheckPendingBinancePaymentTransactionByMoniestUsernameParams struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type CheckPendingBinancePaymentTransactionByMoniestUsernameRow struct {
	Pending       bool      `json:"pending"`
	QrcodeLink    string    `json:"qrcode_link"`
	CheckoutLink  string    `json:"checkout_link"`
	DeepLink      string    `json:"deep_link"`
	UniversalLink string    `json:"universal_link"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) CheckPendingBinancePaymentTransactionByMoniestUsername(ctx context.Context, arg CheckPendingBinancePaymentTransactionByMoniestUsernameParams) ([]CheckPendingBinancePaymentTransactionByMoniestUsernameRow, error) {
	rows, err := q.db.QueryContext(ctx, checkPendingBinancePaymentTransactionByMoniestUsername, arg.UserID, arg.Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CheckPendingBinancePaymentTransactionByMoniestUsernameRow{}
	for rows.Next() {
		var i CheckPendingBinancePaymentTransactionByMoniestUsernameRow
		if err := rows.Scan(
			&i.Pending,
			&i.QrcodeLink,
			&i.CheckoutLink,
			&i.DeepLink,
			&i.UniversalLink,
			&i.CreatedAt,
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

const createBinancePaymentTransactions = `-- name: CreateBinancePaymentTransactions :one
INSERT INTO "binance_payment_transaction" (
        id,
        qrcode_link,
        checkout_link,
        deep_link,
        universal_link,
        status,
        user_id,
        moniest_id,
        date_type,
        date_value,
        moniest_fee,
        amount,
        webhook_url,
        request,
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
        now(),
        now()
    )
RETURNING id, qrcode_link, checkout_link, deep_link, universal_link, status, user_id, moniest_id, date_type, date_value, moniest_fee, amount, webhook_url, payer_id, created_at, updated_at, request, response
`

type CreateBinancePaymentTransactionsParams struct {
	ID            string                 `json:"id"`
	QrcodeLink    string                 `json:"qrcode_link"`
	CheckoutLink  string                 `json:"checkout_link"`
	DeepLink      string                 `json:"deep_link"`
	UniversalLink string                 `json:"universal_link"`
	Status        BinancePaymentStatus   `json:"status"`
	UserID        string                 `json:"user_id"`
	MoniestID     string                 `json:"moniest_id"`
	DateType      BinancePaymentDateType `json:"date_type"`
	DateValue     int32                  `json:"date_value"`
	MoniestFee    float64                `json:"moniest_fee"`
	Amount        float64                `json:"amount"`
	WebhookUrl    string                 `json:"webhook_url"`
	Request       sql.NullString         `json:"request"`
}

func (q *Queries) CreateBinancePaymentTransactions(ctx context.Context, arg CreateBinancePaymentTransactionsParams) (BinancePaymentTransaction, error) {
	row := q.db.QueryRowContext(ctx, createBinancePaymentTransactions,
		arg.ID,
		arg.QrcodeLink,
		arg.CheckoutLink,
		arg.DeepLink,
		arg.UniversalLink,
		arg.Status,
		arg.UserID,
		arg.MoniestID,
		arg.DateType,
		arg.DateValue,
		arg.MoniestFee,
		arg.Amount,
		arg.WebhookUrl,
		arg.Request,
	)
	var i BinancePaymentTransaction
	err := row.Scan(
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
		&i.Request,
		&i.Response,
	)
	return i, err
}

const getBinancePaymentTransaction = `-- name: GetBinancePaymentTransaction :one
SELECT "id",
    "qrcode_link",
    "checkout_link",
    "deep_link",
    "universal_link",
    "status",
    "user_id",
    "moniest_id",
    "date_type",
    "date_value",
    "moniest_fee",
    "amount",
    "webhook_url",
    "payer_id",
    "created_at",
    "updated_at"
FROM "binance_payment_transaction"
WHERE id = $1
`

type GetBinancePaymentTransactionRow struct {
	ID            string                 `json:"id"`
	QrcodeLink    string                 `json:"qrcode_link"`
	CheckoutLink  string                 `json:"checkout_link"`
	DeepLink      string                 `json:"deep_link"`
	UniversalLink string                 `json:"universal_link"`
	Status        BinancePaymentStatus   `json:"status"`
	UserID        string                 `json:"user_id"`
	MoniestID     string                 `json:"moniest_id"`
	DateType      BinancePaymentDateType `json:"date_type"`
	DateValue     int32                  `json:"date_value"`
	MoniestFee    float64                `json:"moniest_fee"`
	Amount        float64                `json:"amount"`
	WebhookUrl    string                 `json:"webhook_url"`
	PayerID       sql.NullString         `json:"payer_id"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

func (q *Queries) GetBinancePaymentTransaction(ctx context.Context, id string) (GetBinancePaymentTransactionRow, error) {
	row := q.db.QueryRowContext(ctx, getBinancePaymentTransaction, id)
	var i GetBinancePaymentTransactionRow
	err := row.Scan(
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
	)
	return i, err
}

const updateBinancePaymentTransactionStatus = `-- name: UpdateBinancePaymentTransactionStatus :one
UPDATE "binance_payment_transaction"
SET "status" = $2,
    "payer_id" = $3,
    "response" = $4,
    updated_at = now()
WHERE id = $1
RETURNING id, qrcode_link, checkout_link, deep_link, universal_link, status, user_id, moniest_id, date_type, date_value, moniest_fee, amount, webhook_url, payer_id, created_at, updated_at, request, response
`

type UpdateBinancePaymentTransactionStatusParams struct {
	ID       string               `json:"id"`
	Status   BinancePaymentStatus `json:"status"`
	PayerID  sql.NullString       `json:"payer_id"`
	Response sql.NullString       `json:"response"`
}

func (q *Queries) UpdateBinancePaymentTransactionStatus(ctx context.Context, arg UpdateBinancePaymentTransactionStatusParams) (BinancePaymentTransaction, error) {
	row := q.db.QueryRowContext(ctx, updateBinancePaymentTransactionStatus,
		arg.ID,
		arg.Status,
		arg.PayerID,
		arg.Response,
	)
	var i BinancePaymentTransaction
	err := row.Scan(
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
		&i.Request,
		&i.Response,
	)
	return i, err
}
