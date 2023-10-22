ALTER TABLE binance_payout_history
ADD "request" TEXT,
    ADD "response" TEXT;

ALTER TABLE binance_payment_transaction
ADD "request" TEXT,
    ADD "response" TEXT;