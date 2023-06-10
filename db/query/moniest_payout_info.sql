-- name: CreateMoniestPayoutInfo :one
INSERT INTO moniest_payout_info (
        id,
        moniest_id,
        source,
        type,
        value,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        now(),
        now()
    )
RETURNING *;