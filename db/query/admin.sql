-- name: UserMetrics :many
SELECT (
        SELECT COUNT(*)
        FROM "user"
    ) AS num_users,
    (
        SELECT COUNT(*)
        FROM moniest
    ) AS num_moniests,
    (
        SELECT COUNT(*)
        FROM user_subscription
        WHERE active = true
    ) AS num_active_subscriptions;

-- name: PostMetrics :many
SELECT COUNT(*) AS num_total_posts,
    COALESCE(
        SUM(
            CASE
                WHEN finished = true THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_finished_posts,
    COALESCE(
        SUM(
            CASE
                WHEN finished = false THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_unfinished_posts,
    COALESCE(
        SUM(
            CASE
                WHEN status = 'success' THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_success_posts,
    COALESCE(
        SUM(
            CASE
                WHEN status = 'fail' THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_fail_posts,
    COALESCE(
        SUM(
            CASE
                WHEN status = 'pending' THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_pending_posts,
    COALESCE(
        SUM(
            CASE
                WHEN market_type = 'futures' THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_futures_posts,
    COALESCE(
        SUM(
            CASE
                WHEN market_type = 'spot' THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_spot_posts,
    COALESCE(COUNT(DISTINCT moniest_id), 0) AS num_unique_moniests
FROM post_crypto;

-- name: PaymentMetrics :many
SELECT SUM(
        CASE
            WHEN status = 'success' THEN 1
            ELSE 0
        END
    ) AS num_success_payment,
    SUM(
        CASE
            WHEN status = 'success' THEN amount
            ELSE 0
        END
    )::double precision AS success_payment_amount,
    SUM(
        CASE
            WHEN status = 'fail' THEN 1
            ELSE 0
        END
    ) AS num_fail_payment,
    SUM(
        CASE
            WHEN status = 'fail' THEN amount
            ELSE 0
        END
    )::double precision AS fail_payment_amount,
    SUM(
        CASE
            WHEN status = 'pending' THEN 1
            ELSE 0
        END
    ) AS num_pending_payment,
    SUM(
        CASE
            WHEN status = 'pending' THEN amount
            ELSE 0
        END
    )::double precision AS pending_payment_amount
FROM binance_payment_transaction;

-- name: PayoutMetrics :many
SELECT COUNT(*) AS num_payouts,
    COUNT(DISTINCT user_id) AS num_unique_users,
    COUNT(DISTINCT moniest_id) AS num_unique_moniests,
    COUNT(DISTINCT transaction_id) AS num_unique_transactions,
    -- success payouts
    SUM(
        CASE
            WHEN status = 'success' THEN 1
            ELSE 0
        END
    ) AS num_success_payouts,
    SUM(
        CASE
            WHEN status = 'success' THEN amount
            ELSE 0
        END
    )::double precision AS success_payouts_amount,
    SUM(
        CASE
            WHEN status = 'success' THEN amount * (
                1 - (COALESCE(operation_fee_percentage, $1) / 100)
            ) -- Apply the cut percentage
            ELSE 0
        END
    )::double precision AS success_payouts_amount_after_cut,
    -- fail payouts
    SUM(
        CASE
            WHEN status = 'fail' THEN 1
            ELSE 0
        END
    ) AS num_fail_payouts,
    SUM(
        CASE
            WHEN status = 'fail' THEN amount
            ELSE 0
        END
    )::double precision AS fail_payouts_amount,
    -- pending payouts
    SUM(
        CASE
            WHEN status = 'pending' THEN 1
            ELSE 0
        END
    ) AS num_pending_payouts,
    SUM(
        CASE
            WHEN status = 'pending' THEN amount
            ELSE 0
        END
    )::double precision AS pending_payouts_amount,
    SUM(
        CASE
            WHEN status = 'pending' THEN amount * (
                1 - (COALESCE(operation_fee_percentage, $1) / 100)
            ) -- Apply the cut percentage
            ELSE 0
        END
    )::double precision AS pending_payouts_amount_after_cut,
    -- redund payouts
    SUM(
        CASE
            WHEN status = 'refund' THEN 1
            ELSE 0
        END
    ) AS num_refund_payouts,
    SUM(
        CASE
            WHEN status = 'refund' THEN amount
            ELSE 0
        END
    )::double precision AS refund_payouts_amount,
    SUM(
        CASE
            WHEN status = 'refund' THEN amount * (
                1 - (COALESCE(operation_fee_percentage, $1) / 100)
            ) -- Apply the cut percentage
            ELSE 0
        END
    )::double precision AS refund_payouts_amount_after_cut,
    -- redund fail payouts
    SUM(
        CASE
            WHEN status = 'refund_fail' THEN 1
            ELSE 0
        END
    ) AS num_refund_fail_payouts,
    SUM(
        CASE
            WHEN status = 'refund_fail' THEN amount
            ELSE 0
        END
    )::double precision AS refund_fail_payouts_amount
FROM binance_payout_history;

-- name: FeedbackMetrics :many
SELECT COALESCE(COUNT(*), 0) AS num_all_feedbacks,
    COALESCE(
        SUM(
            CASE
                WHEN solved = true THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_solved_feedbacks,
    COALESCE(
        SUM(
            CASE
                WHEN solved = false THEN 1
                ELSE 0
            END
        ),
        0
    ) AS num_unsolved_feedbacks
FROM feedback;

-- name: GetFeedbacks :many
SELECT id,
    COALESCE(user_id, '') AS user_id,
    COALESCE(type, '') AS type,
    message,
    solved,
    created_at
FROM feedback;

-- name: ADMIN_GetAllUsers :many
SELECT "user"."id",
    "moniest"."id" as "moniest_id",
    "user"."fullname",
    "user"."username",
    "user"."email",
    "user"."email_verified",
    "user"."language",
    "user"."location",
    "user"."created_at",
    "user"."updated_at",
    "moniest"."bio",
    "moniest"."description",
    "mpcs"."pnl_7days",
    "mpcs"."roi_7days",
    "mpcs"."win_rate_7days",
    "mpcs"."pnl_30days",
    "mpcs"."roi_30days",
    "mpcs"."win_rate_30days",
    "mpcs"."pnl_total",
    "mpcs"."roi_total",
    "mpcs"."win_rate_total",
    "moniest_subscription_info"."id" as "moniest_subscription_info_id",
    "moniest_subscription_info"."fee",
    "moniest_subscription_info"."message",
    "moniest_subscription_info"."updated_at" as "moniest_subscription_info_updated_at",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "user"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "user"."id"
                AND "image"."type" = 'profile_photo'
        ),
        ''
    ) AS "profile_photo_thumbnail_link",
    COALESCE (
        (
            SELECT "image"."link"
            FROM "image"
            WHERE "image"."user_id" = "user"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_link",
    COALESCE (
        (
            SELECT "image"."thumbnail_link"
            FROM "image"
            WHERE "image"."user_id" = "user"."id"
                AND "image"."type" = 'background_photo'
        ),
        ''
    ) AS "background_photo_thumbnail_link"
FROM "user"
    LEFT JOIN "moniest" ON "moniest"."user_id" = "user"."id"
    LEFT JOIN "moniest_subscription_info" ON "moniest_subscription_info"."moniest_id" = "moniest"."id"
    LEFT JOIN "moniest_post_crypto_statistics" AS mpcs ON "mpcs"."moniest_id" = "moniest"."id"
ORDER BY "user"."created_at" DESC
LIMIT $1 OFFSET $2;