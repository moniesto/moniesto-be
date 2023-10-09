// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: admin.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const feedbackMetrics = `-- name: FeedbackMetrics :many
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
FROM feedback
`

type FeedbackMetricsRow struct {
	NumAllFeedbacks      interface{} `json:"num_all_feedbacks"`
	NumSolvedFeedbacks   interface{} `json:"num_solved_feedbacks"`
	NumUnsolvedFeedbacks interface{} `json:"num_unsolved_feedbacks"`
}

func (q *Queries) FeedbackMetrics(ctx context.Context) ([]FeedbackMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, feedbackMetrics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FeedbackMetricsRow{}
	for rows.Next() {
		var i FeedbackMetricsRow
		if err := rows.Scan(&i.NumAllFeedbacks, &i.NumSolvedFeedbacks, &i.NumUnsolvedFeedbacks); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedbacks = `-- name: GetFeedbacks :many
SELECT id,
    COALESCE(user_id, '') AS user_id,
    COALESCE(type, '') AS type,
    message,
    solved,
    created_at
FROM feedback
`

type GetFeedbacksRow struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	Solved    bool      `json:"solved"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) GetFeedbacks(ctx context.Context) ([]GetFeedbacksRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedbacks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFeedbacksRow{}
	for rows.Next() {
		var i GetFeedbacksRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Type,
			&i.Message,
			&i.Solved,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const paymentMetrics = `-- name: PaymentMetrics :many
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
FROM binance_payment_transaction
`

type PaymentMetricsRow struct {
	NumSuccessPayment    int64   `json:"num_success_payment"`
	SuccessPaymentAmount float64 `json:"success_payment_amount"`
	NumFailPayment       int64   `json:"num_fail_payment"`
	FailPaymentAmount    float64 `json:"fail_payment_amount"`
	NumPendingPayment    int64   `json:"num_pending_payment"`
	PendingPaymentAmount float64 `json:"pending_payment_amount"`
}

func (q *Queries) PaymentMetrics(ctx context.Context) ([]PaymentMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, paymentMetrics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PaymentMetricsRow{}
	for rows.Next() {
		var i PaymentMetricsRow
		if err := rows.Scan(
			&i.NumSuccessPayment,
			&i.SuccessPaymentAmount,
			&i.NumFailPayment,
			&i.FailPaymentAmount,
			&i.NumPendingPayment,
			&i.PendingPaymentAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const payoutMetrics = `-- name: PayoutMetrics :many
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
FROM binance_payout_history
`

type PayoutMetricsRow struct {
	NumPayouts                   int64   `json:"num_payouts"`
	NumUniqueUsers               int64   `json:"num_unique_users"`
	NumUniqueMoniests            int64   `json:"num_unique_moniests"`
	NumUniqueTransactions        int64   `json:"num_unique_transactions"`
	NumSuccessPayouts            int64   `json:"num_success_payouts"`
	SuccessPayoutsAmount         float64 `json:"success_payouts_amount"`
	SuccessPayoutsAmountAfterCut float64 `json:"success_payouts_amount_after_cut"`
	NumFailPayouts               int64   `json:"num_fail_payouts"`
	FailPayoutsAmount            float64 `json:"fail_payouts_amount"`
	NumPendingPayouts            int64   `json:"num_pending_payouts"`
	PendingPayoutsAmount         float64 `json:"pending_payouts_amount"`
	PendingPayoutsAmountAfterCut float64 `json:"pending_payouts_amount_after_cut"`
	NumRefundPayouts             int64   `json:"num_refund_payouts"`
	RefundPayoutsAmount          float64 `json:"refund_payouts_amount"`
	RefundPayoutsAmountAfterCut  float64 `json:"refund_payouts_amount_after_cut"`
	NumRefundFailPayouts         int64   `json:"num_refund_fail_payouts"`
	RefundFailPayoutsAmount      float64 `json:"refund_fail_payouts_amount"`
}

func (q *Queries) PayoutMetrics(ctx context.Context, operationFeePercentage sql.NullFloat64) ([]PayoutMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, payoutMetrics, operationFeePercentage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PayoutMetricsRow{}
	for rows.Next() {
		var i PayoutMetricsRow
		if err := rows.Scan(
			&i.NumPayouts,
			&i.NumUniqueUsers,
			&i.NumUniqueMoniests,
			&i.NumUniqueTransactions,
			&i.NumSuccessPayouts,
			&i.SuccessPayoutsAmount,
			&i.SuccessPayoutsAmountAfterCut,
			&i.NumFailPayouts,
			&i.FailPayoutsAmount,
			&i.NumPendingPayouts,
			&i.PendingPayoutsAmount,
			&i.PendingPayoutsAmountAfterCut,
			&i.NumRefundPayouts,
			&i.RefundPayoutsAmount,
			&i.RefundPayoutsAmountAfterCut,
			&i.NumRefundFailPayouts,
			&i.RefundFailPayoutsAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const postMetrics = `-- name: PostMetrics :many
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
FROM post_crypto
`

type PostMetricsRow struct {
	NumTotalPosts      int64 `json:"num_total_posts"`
	NumFinishedPosts   int64 `json:"num_finished_posts"`
	NumUnfinishedPosts int64 `json:"num_unfinished_posts"`
	NumSuccessPosts    int64 `json:"num_success_posts"`
	NumFailPosts       int64 `json:"num_fail_posts"`
	NumPendingPosts    int64 `json:"num_pending_posts"`
	NumFuturesPosts    int64 `json:"num_futures_posts"`
	NumSpotPosts       int64 `json:"num_spot_posts"`
	NumUniqueMoniests  int64 `json:"num_unique_moniests"`
}

func (q *Queries) PostMetrics(ctx context.Context) ([]PostMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, postMetrics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PostMetricsRow{}
	for rows.Next() {
		var i PostMetricsRow
		if err := rows.Scan(
			&i.NumTotalPosts,
			&i.NumFinishedPosts,
			&i.NumUnfinishedPosts,
			&i.NumSuccessPosts,
			&i.NumFailPosts,
			&i.NumPendingPosts,
			&i.NumFuturesPosts,
			&i.NumSpotPosts,
			&i.NumUniqueMoniests,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userMetrics = `-- name: UserMetrics :many
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
    ) AS num_active_subscriptions
`

type UserMetricsRow struct {
	NumUsers               int64 `json:"num_users"`
	NumMoniests            int64 `json:"num_moniests"`
	NumActiveSubscriptions int64 `json:"num_active_subscriptions"`
}

func (q *Queries) UserMetrics(ctx context.Context) ([]UserMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, userMetrics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserMetricsRow{}
	for rows.Next() {
		var i UserMetricsRow
		if err := rows.Scan(&i.NumUsers, &i.NumMoniests, &i.NumActiveSubscriptions); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
