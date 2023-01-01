-- name: CreatePost :one
INSERT INTO "post_crypto" (
        id,
        moniest_id,
        currency,
        start_price,
        duration,
        target1,
        target2,
        target3,
        stop,
        direction,
        score,
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
        now(),
        now()
    )
RETURNING id,
    moniest_id,
    currency,
    start_price,
    duration,
    target1,
    target2,
    target3,
    stop,
    direction,
    score,
    created_at,
    updated_at;

-- -- name: DeletePost :one
-- UPDATE "post_crypto"
-- SET "deleted" = true,
--     "updated_at" = now()
-- WHERE "id" = $1
-- RETURNING *;
-- -- TODO:
-- -- Get moniests live posts, ended posts, all posts by username
-- -- name: GetActivePostsByUsername :many
-- SELECT "post_crypto".*,
--     "post_crypto_description"."description"
-- FROM "post_crypto"
--     INNER JOIN "moniest" ON "moniest"."id" = "post_crypto"."moniest_id"
--     INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
--     INNER JOIN "post_crypto_description" ON "post_crypto_description"."post_id" = "post_crypto"."id"
-- WHERE "user"."username" = $1
--     AND "user"."deleted" = false
--     AND duration > now();
-- -- name: GetInactivePostsByUsername :many
-- SELECT "post_crypto".*,
--     "post_crypto_description"."description"
-- FROM "post_crypto"
--     INNER JOIN "moniest" ON "moniest"."id" = "post_crypto"."moniest_id"
--     INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
--     INNER JOIN "post_crypto_description" ON "post_crypto_description"."post_id" = "post_crypto"."id"
-- WHERE "user"."username" = $1
--     AND "user"."deleted" = false
--     AND duration < now();
-- -- name: GetAllPostsByUsername :many
-- SELECT "post_crypto".*,
--     "post_crypto_description"."description"
-- FROM "post_crypto"
--     INNER JOIN "moniest" ON "moniest"."id" = "post_crypto"."moniest_id"
--     INNER JOIN "user" ON "user"."id" = "moniest"."user_id"
--     INNER JOIN "post_crypto_description" ON "post_crypto_description"."post_id" = "post_crypto"."id"
-- WHERE "user"."username" = $1
--     AND "user"."deleted" = false;