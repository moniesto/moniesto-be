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
    "last_job_timestamp" = $2
WHERE "id" = $3;

-- name: UpdateFinishedPostStatus :exec
UPDATE "post_crypto"
SET "status" = $1,
    "score" = $2,
    "finished" = TRUE
WHERE "id" = $3;

-- name: UpdateMoniestScore :exec
UPDATE "moniest"
SET "score" = GREATEST("score" + $1, 0)
WHERE "id" = $2;