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

-- name: GetMoniestPayoutInfos :many
SELECT "id",
    "moniest_id",
    "source",
    "type",
    "value",
    "updated_at"
FROM "moniest_payout_info"
WHERE "moniest_id" = $1;

-- name: UpdateMoniestPayoutInfo :exec
UPDATE "moniest_payout_info"
SET "value" = $4,
    "updated_at" = now()
WHERE "moniest_id" = $1
    AND "source" = $2
    AND "type" = $3;