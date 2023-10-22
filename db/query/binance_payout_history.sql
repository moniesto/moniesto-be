-- name: CreateBinancePayoutHistory :one
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
RETURNING *;

-- name: GetBinancePayoutHistories :many
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
ORDER BY date_index ASC;

-- name: UpdateBinancePayoutHistoryRefund :exec
UPDATE "binance_payout_history"
SET "status" = $2,
    "failure_message" = $3,
    "request" = $4,
    "response" = $5,
    updated_at = now()
WHERE id = $1;