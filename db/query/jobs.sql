-- name: GetAllActivePosts :many
SELECT "pc"."id",
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
WHERE "pc"."duration" > now()
    AND "pc"."finished" = FALSE
ORDER BY "pc"."created_at" ASC;