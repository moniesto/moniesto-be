-- name: CreateMoniestPostCryptoStatistics :one
INSERT INTO moniest_post_crypto_statistics (
        "id",
        "moniest_id",
        "pnl_7days",
        "roi_7days",
        "win_rate_7days",
        "posts_7days",
        "pnl_30days",
        "roi_30days",
        "win_rate_30days",
        "posts_30days",
        "pnl_total",
        "roi_total",
        "win_rate_total",
        "created_at"
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
        now()
    )
RETURNING *;

-- name: UpdateAllMoniestsPostCryptoStatistics_7days :exec
UPDATE moniest_post_crypto_statistics AS mpcs
SET pnl_7days = COALESCE(
        (
            SELECT ROUND(SUM(pnl)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        0
    ),
    roi_7days = COALESCE(
        (
            SELECT ROUND(AVG(roi)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        0
    ),
    win_rate_7days = COALESCE(
        (
            SELECT ROUND(
                    (
                        (
                            SUM(
                                CASE
                                    WHEN pc.status = 'success' THEN 1
                                    ELSE 0
                                END
                            )::float / COUNT(*)
                        ) * 100
                    )::numeric,
                    2
                )
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        0
    ),
    posts_7days = COALESCE(
        (
            SELECT ARRAY_AGG(id)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        '{}'
    ),
    updated_at = now();

-- name: UpdateAllMoniestsPostCryptoStatistics_30days :exec
UPDATE moniest_post_crypto_statistics AS mpcs
SET pnl_30days = COALESCE(
        (
            SELECT ROUND(SUM(pnl)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        0
    ),
    roi_30days = COALESCE(
        (
            SELECT ROUND(AVG(roi)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        0
    ),
    win_rate_30days = COALESCE(
        (
            SELECT ROUND(
                    (
                        (
                            SUM(
                                CASE
                                    WHEN pc.status = 'success' THEN 1
                                    ELSE 0
                                END
                            )::float / COUNT(*)
                        ) * 100
                    )::numeric,
                    2
                )
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        0
    ),
    posts_30days = COALESCE(
        (
            SELECT ARRAY_AGG(id)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        '{}'
    ),
    updated_at = now();

-- name: UpdateAllMoniestsPostCryptoStatistics_total :exec
UPDATE moniest_post_crypto_statistics AS mpcs
SET pnl_total = COALESCE(
        (
            SELECT ROUND(SUM(pnl)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
        ),
        0
    ),
    roi_total = COALESCE(
        (
            SELECT ROUND(AVG(roi)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
        ),
        0
    ),
    win_rate_total = COALESCE(
        (
            SELECT ROUND(
                    (
                        (
                            SUM(
                                CASE
                                    WHEN pc.status = 'success' THEN 1
                                    ELSE 0
                                END
                            )::float / COUNT(*)
                        ) * 100
                    )::numeric,
                    2
                )
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
        ),
        0
    ),
    updated_at = now();

-- name: UpdateMoniestsPostCryptoStatistics_7days :exec
UPDATE moniest_post_crypto_statistics AS mpcs
SET pnl_7days = COALESCE(
        (
            SELECT ROUND(SUM(pnl)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        0
    ),
    roi_7days = COALESCE(
        (
            SELECT ROUND(AVG(roi)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        0
    ),
    win_rate_7days = COALESCE(
        (
            SELECT ROUND(
                    (
                        (
                            SUM(
                                CASE
                                    WHEN pc.status = 'success' THEN 1
                                    ELSE 0
                                END
                            )::float / COUNT(*)
                        ) * 100
                    )::numeric,
                    2
                )
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        0
    ),
    posts_7days = COALESCE(
        (
            SELECT ARRAY_AGG(id)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = TRUE
                AND pc.finished_at >= NOW() - INTERVAL '7 days'
        ),
        '{}'
    ),
    updated_at = now()
WHERE "mpcs"."moniest_id" = ANY($1::varchar []);

-- name: UpdateMoniestsPostCryptoStatistics_30days :exec
UPDATE moniest_post_crypto_statistics AS mpcs
SET pnl_30days = COALESCE(
        (
            SELECT ROUND(SUM(pnl)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        0
    ),
    roi_30days = COALESCE(
        (
            SELECT ROUND(AVG(roi)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        0
    ),
    win_rate_30days = COALESCE(
        (
            SELECT ROUND(
                    (
                        (
                            SUM(
                                CASE
                                    WHEN pc.status = 'success' THEN 1
                                    ELSE 0
                                END
                            )::float / COUNT(*)
                        ) * 100
                    )::numeric,
                    2
                )
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        0
    ),
    posts_30days = COALESCE(
        (
            SELECT ARRAY_AGG(id)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
                AND pc.finished_at >= NOW() - INTERVAL '30 days'
        ),
        '{}'
    ),
    updated_at = now()
WHERE "mpcs"."moniest_id" = ANY($1::varchar []);

-- name: UpdateMoniestsPostCryptoStatistics_total :exec
UPDATE moniest_post_crypto_statistics AS mpcs
SET pnl_total = COALESCE(
        (
            SELECT ROUND(SUM(pnl)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
        ),
        0
    ),
    roi_total = COALESCE(
        (
            SELECT ROUND(AVG(roi)::numeric, 2)
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
        ),
        0
    ),
    win_rate_total = COALESCE(
        (
            SELECT ROUND(
                    (
                        (
                            SUM(
                                CASE
                                    WHEN pc.status = 'success' THEN 1
                                    ELSE 0
                                END
                            )::float / COUNT(*)
                        ) * 100
                    )::numeric,
                    2
                )
            FROM post_crypto AS pc
            WHERE pc.moniest_id = mpcs.moniest_id
                AND pc.finished = true
        ),
        0
    ),
    updated_at = now()
WHERE "mpcs"."moniest_id" = ANY($1::varchar []);