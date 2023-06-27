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