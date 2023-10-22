ALTER TABLE binance_payout_history DROP COLUMN "request",
    DROP COLUMN "response";

ALTER TABLE binance_payment_transaction DROP COLUMN "request",
    DROP COLUMN "response";