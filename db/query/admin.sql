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
    SUM(
        CASE
            WHEN finished = true THEN 1
            ELSE 0
        END
    ) AS num_finished_posts,
    SUM(
        CASE
            WHEN finished = false THEN 1
            ELSE 0
        END
    ) AS num_unfinished_posts,
    SUM(
        CASE
            WHEN status = 'success' THEN 1
            ELSE 0
        END
    ) AS num_success_posts,
    SUM(
        CASE
            WHEN status = 'fail' THEN 1
            ELSE 0
        END
    ) AS num_fail_posts,
    SUM(
        CASE
            WHEN status = 'pending' THEN 1
            ELSE 0
        END
    ) AS num_pending_posts,
    SUM(
        CASE
            WHEN market_type = 'futures' THEN 1
            ELSE 0
        END
    ) AS num_futures_posts,
    SUM(
        CASE
            WHEN market_type = 'spot' THEN 1
            ELSE 0
        END
    ) AS num_spot_posts,
    COUNT(DISTINCT moniest_id) AS num_unique_moniests
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
    ) AS success_payment_amount,
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
    ) AS fail_payment_amount,
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
    ) AS pending_payment_amount
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
    ) AS success_payouts_amount,
    SUM(
        CASE
            WHEN status = 'success' THEN amount * (
                1 - (COALESCE(operation_fee_percentage, 18) / 100)
            ) -- Apply the cut percentage
            ELSE 0
        END
    ) AS success_payouts_amount_after_cut,
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
    ) AS fail_payouts_amount,
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
    ) AS pending_payouts_amount,
    SUM(
        CASE
            WHEN status = 'pending' THEN amount * (
                1 - (COALESCE(operation_fee_percentage, 18) / 100)
            ) -- Apply the cut percentage
            ELSE 0
        END
    ) AS pending_payouts_amount_after_cut,
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
    ) AS refund_payouts_amount,
    SUM(
        CASE
            WHEN status = 'refund' THEN amount * (
                1 - (COALESCE(operation_fee_percentage, 18) / 100)
            ) -- Apply the cut percentage
            ELSE 0
        END
    ) AS refund_payouts_amount_after_cut,
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
    ) AS refund_fail_payouts_amount
FROM binance_payout_history;