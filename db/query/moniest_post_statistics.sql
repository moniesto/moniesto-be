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