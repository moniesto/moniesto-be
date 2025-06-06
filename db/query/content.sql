-- name: GetSubscribedActivePosts :many
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
    INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
    AND "us"."user_id" = $1
    AND "us"."active" = TRUE
    AND "pc"."duration" > now()
    AND "pc"."finished" = FALSE
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetPostByID :one
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
    "m"."id" AS "moniest_id",
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
    "u"."id" AS "user_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "pcd"."description" AS "post_description",
    COALESCE(
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE(
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE(
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE(
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "u"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "post_crypto" AS pc
    INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
    INNER JOIN "moniest" AS m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "user" AS u ON "m"."user_id" = "u"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    LEFT JOIN "post_crypto_description" AS pcd ON "pcd"."post_id" = "pc"."id"
WHERE "pc"."id" = $1
LIMIT 1;

-- name: GetSubscribedActivePostsWithOwn :many
(
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
        INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
        AND "us"."user_id" = $1
        AND "us"."active" = TRUE
        AND "pc"."duration" > now()
        AND "pc"."finished" = FALSE
        INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
        INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
        INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
        LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
)
UNION ALL
(
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
        AND "m"."user_id" = $1
        AND "pc"."duration" > now()
        AND "pc"."finished" = FALSE
        INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
        INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
        LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
)
ORDER BY "created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetSubscribedDeactivePosts :many
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
    INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
    AND "us"."user_id" = $1
    AND "us"."active" = TRUE
    AND (
        "pc"."duration" < now()
        OR "pc"."finished" = TRUE
    )
    INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
    INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
    INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetSubscribedDeactivePostsWithOwn :many
(
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
        INNER JOIN "user_subscription" AS us ON "pc"."moniest_id" = "us"."moniest_id"
        AND "us"."user_id" = $1
        AND "us"."active" = TRUE
        AND (
            "pc"."duration" < now()
            OR "pc"."finished" = TRUE
        )
        INNER JOIN "moniest" as m ON "pc"."moniest_id" = "m"."id"
        INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
        INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
        LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
)
UNION ALL
(
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
        AND "m"."user_id" = $1
        AND (
            "pc"."duration" < now()
            OR "pc"."finished" = TRUE
        )
        INNER JOIN "user" as u ON "m"."user_id" = "u"."id"
        INNER JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "m"."id"
        LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
)
ORDER BY "created_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetDeactivePostsByPNL :many
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
    AND (
        "pc"."duration" < now()
        OR "pc"."finished" = TRUE
    )
    AND "pc"."status" = 'success'
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."pnl" DESC
LIMIT $1 OFFSET $2;

-- name: GetDeactivePostsByCreatedAt :many
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
    AND (
        "pc"."duration" < now()
        OR "pc"."finished" = TRUE
    )
    AND "pc"."status" = 'success'
    LEFT JOIN "post_crypto_description" as pcd ON "pcd"."post_id" = "pc"."id"
ORDER BY "pc"."created_at" DESC
LIMIT $1 OFFSET $2;

-- name: GetMoniests :many
SELECT "u"."id",
    "m"."id" as "moniest_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "u"."created_at",
    "u"."updated_at",
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
    "msi"."fee",
    "msi"."message",
    "msi"."updated_at" as "moniest_subscription_info_updated_at",
    COUNT("us"."id") as "user_subscription_count",
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
FROM "moniest" as m
    INNER JOIN "user" as u ON "u"."id" = "m"."user_id"
    INNER JOIN "moniest_subscription_info" as msi ON "msi"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" as mpcs ON "mpcs"."moniest_id" = "m"."id"
    LEFT JOIN "user_subscription" as us on "us"."moniest_id" = "m"."id"
    AND "us"."active" = TRUE
    AND "u"."deleted" = FALSE
GROUP BY "u"."id",
    "m"."id",
    "msi"."id",
    "mpcs"."id"
ORDER BY "mpcs"."pnl_30days" DESC,
    "u"."created_at"
LIMIT $1 OFFSET $2;

-- name: SearchMoniests :many
SELECT "u"."id",
    "m"."id" as "moniest_id",
    "u"."fullname",
    "u"."username",
    "u"."email_verified",
    "u"."location",
    "u"."created_at",
    "u"."updated_at",
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
    "msi"."fee",
    "msi"."message",
    "msi"."updated_at" as "moniest_subscription_info_updated_at",
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
FROM "user" as "u"
    INNER JOIN "moniest" as "m" ON "m"."user_id" = "u"."id"
    INNER JOIN "moniest_subscription_info" as "msi" ON "msi"."moniest_id" = "m"."id"
    INNER JOIN "moniest_post_crypto_statistics" as mpcs ON "mpcs"."moniest_id" = "m"."id"
WHERE ("u"."fullname" ILIKE $1)
    OR ("u"."username" ILIKE $1)
    AND "u"."deleted" = FALSE
LIMIT $2 OFFSET $3;