-- name: CreateBinancePaymentTransactions :one
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
        now(),
        now()
    )
RETURNING *;