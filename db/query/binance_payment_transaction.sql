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

-- name: GetBinancePaymentTransaction :one
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
WHERE id = $1;

-- name: UpdateBinancePaymentTransactionStatus :one
UPDATE "binance_payment_transaction"
SET "status" = $2,
    "payer_id" = $3,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: CheckPendingBinancePaymentTransactionByMoniestUsername :one
SELECT COUNT(*) != 0 as pending
FROM "binance_payment_transaction"
    INNER JOIN "moniest" ON "moniest"."id" = "binance_payment_transaction"."moniest_id"
    INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
    AND "user"."username" = $2
WHERE "binance_payment_transaction"."status" = 'pending'
    AND "binance_payment_transaction"."user_id" = $1;