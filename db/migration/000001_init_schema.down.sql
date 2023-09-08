-- drop tables
DROP TABLE IF EXISTS "binance_payout_history";

DROP TABLE IF EXISTS "binance_payment_transaction";

DROP TABLE IF EXISTS "feedback";

DROP TABLE IF EXISTS "email_verification_token";

DROP TABLE IF EXISTS "password_reset_token";

DROP TABLE IF EXISTS "post_crypto_description";

DROP TABLE IF EXISTS "post_crypto";

DROP TABLE IF EXISTS "user_subscription_history";

DROP TABLE IF EXISTS "user_subscription";

DROP TABLE IF EXISTS "moniest_subscription_info";

DROP TABLE IF EXISTS "moniest_payout_info";

DROP TABLE IF EXISTS "moniest_post_crypto_statistics";

DROP TABLE IF EXISTS "moniest";

DROP TABLE IF EXISTS "image";

DROP TABLE IF EXISTS "user";

-- DROP TABLE IF EXISTS "schema_migrations";
-- 
-- drop types
DROP TYPE IF EXISTS "payout_type";

DROP TYPE IF EXISTS "payout_source";

DROP TYPE IF EXISTS "binance_payment_date_type";

DROP TYPE IF EXISTS "binance_payout_status";

DROP TYPE IF EXISTS "binance_payment_status";

DROP TYPE IF EXISTS "post_crypto_status";

DROP TYPE IF EXISTS "entry_position";

DROP TYPE IF EXISTS "post_crypto_market_type";

DROP TYPE IF EXISTS "image_type";

DROP TYPE IF EXISTS "user_language";