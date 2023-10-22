-- name: GetAllActivePosts :many
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
ORDER BY "pc"."created_at" ASC;

-- name: UpdateUnfinishedPostStatus :exec
UPDATE "post_crypto"
SET "last_operated_at" = $2,
    updated_at = now()
WHERE "id" = $1;

-- name: UpdateFinishedPostStatus :exec
UPDATE "post_crypto"
SET "status" = $2,
    "pnl" = $3,
    "roi" = $4,
    "hit_price" = $5,
    "last_operated_at" = $6,
    "finished" = TRUE,
    "finished_at" = now(),
    updated_at = now()
WHERE "id" = $1;

-- name: GetAllPendingPayouts :many
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
    AND payout_date <= now();

-- name: UpdateBinancePayoutHistoryPayout :exec
UPDATE "binance_payout_history"
SET "status" = $2,
    operation_fee_percentage = $3,
    "payout_done_at" = $4,
    "failure_message" = $5,
    payout_request_id = $6,
    "request" = $7,
    "response" = $8,
    updated_at = now()
WHERE "id" = $1;

-- name: GetExpiredActiveSubscriptions :many
SELECT *
FROM "user_subscription"
WHERE active = TRUE
    AND subscription_end_date <= now();

-- name: UpdateExpiredActiveSubscription :exec
UPDATE "user_subscription"
SET active = FALSE,
    updated_at = now()
WHERE "id" = $1;

-- name: CreateUserSubscriptionHistory :one
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
RETURNING *;

-- name: GetExpiredPendingBinanceTransactions :many
SELECT *
FROM binance_payment_transaction
WHERE status = 'pending'
    AND "created_at" + INTERVAL '5 minutes' <= NOW();

-- name: UpdateExpiredPendingBinanceTransaction :exec
UPDATE "binance_payment_transaction"
SET status = 'fail',
    updated_at = now()
WHERE "id" = $1;