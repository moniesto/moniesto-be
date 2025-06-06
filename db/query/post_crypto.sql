-- name: CreatePost :one
INSERT INTO "post_crypto" (
        id,
        moniest_id,
        market_type,
        currency,
        start_price,
        duration,
        take_profit,
        stop,
        target1,
        target2,
        target3,
        direction,
        leverage,
        pnl,
        roi,
        last_operated_at,
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
        $16,
        now(),
        now()
    )
RETURNING id,
    moniest_id,
    market_type,
    currency,
    start_price,
    duration,
    take_profit,
    stop,
    target1,
    target2,
    target3,
    direction,
    leverage,
    pnl,
    roi,
    created_at,
    updated_at;

-- name: GetMoniestAllPostsByUsername :many
SELECT "pc"."id",
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
    "pc"."hit_price",
    "pc"."finished_at",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "u"."id" as "user_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "pcd"."description" as "post_description",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "post_crypto" AS pc
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    AND "u"."username" = $1
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetMoniestActivePostsByUsername :many
SELECT "pc"."id",
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
    "pc"."hit_price",
    "pc"."finished_at",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "u"."id" as "user_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "pcd"."description" as "post_description",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "post_crypto" AS pc
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    AND "u"."username" = $1
    AND "pc"."duration" > now()
    AND "pc"."finished" = FALSE
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetMoniestDeactivePostsByUsername :many
SELECT "pc"."id",
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
    "pc"."hit_price",
    "pc"."finished_at",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "u"."id" as "user_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "pcd"."description" as "post_description",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "post_crypto" AS pc
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    AND "u"."username" = $1
    AND (
        "pc"."duration" < now()
        OR "pc"."finished" = TRUE
    )
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetOwnAllPostsByUsername :many
SELECT "pc"."id",
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
    "pc"."hit_price",
    "pc"."finished_at",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "u"."id" as "user_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "pcd"."description" as "post_description",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "post_crypto" AS pc
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    AND "u"."username" = $1
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetOwnActivePostsByUsername :many
SELECT "pc"."id",
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
    "pc"."hit_price",
    "pc"."finished_at",
    "pc"."created_at",
    "pc"."updated_at",
    "m"."id" as "moniest_id",
    "m"."bio",
    "m"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "u"."id" as "user_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "pcd"."description" as "post_description",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "post_crypto" AS pc
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    AND "u"."username" = $1
    AND "pc"."duration" > now()
    AND "pc"."finished" = FALSE
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

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