-- name: GetAllActivePosts :many
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
WHERE "pc"."finished" = FALSE
ORDER BY "pc"."created_at" ASC;

-- name: UpdateUnfinishedPostStatus :exec
UPDATE "post_crypto"
SET "last_target_hit" = $1,
    "last_job_timestamp" = $2,
    updated_at = now()
WHERE "id" = $3;

-- name: UpdateFinishedPostStatus :exec
UPDATE "post_crypto"
SET "status" = $1,
    "score" = $2,
    "finished" = TRUE,
    updated_at = now()
WHERE "id" = $3;

-- name: UpdateMoniestScore :exec
UPDATE "moniest"
SET "score" = GREATEST("score" + $1, 0),
    updated_at = now()
WHERE "id" = $2;

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

-- name: UpdatePayoutHistory :exec
UPDATE "binance_payout_history"
SET "status" = $2,
    operation_fee_percentage = $3,
    "payout_done_at" = $4,
    "failure_message" = $5,
    payout_request_id = $6,
    updated_at = now()
WHERE "id" = $1;