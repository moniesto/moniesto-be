CREATE TYPE "user_language" AS ENUM ('en', 'tr');

CREATE TYPE "image_type" AS ENUM ('profile_photo', 'background_photo');

CREATE TYPE "post_crypto_market_type" AS ENUM ('spot', 'futures');

CREATE TYPE "direction" AS ENUM ('long', 'short');

CREATE TYPE "post_crypto_status" AS ENUM ('pending', 'fail', 'success');

CREATE TYPE "binance_payment_status" AS ENUM ('pending', 'fail', 'success');

CREATE TYPE "binance_payout_status" AS ENUM (
  'pending',
  'fail',
  'success',
  'refund',
  'refund_fail'
);

CREATE TYPE "binance_payment_date_type" AS ENUM ('MONTH');

CREATE TYPE "payout_source" AS ENUM ('BINANCE');

CREATE TYPE "payout_type" AS ENUM ('BINANCE_ID');

CREATE TABLE "user" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "fullname" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "email_verified" boolean NOT NULL DEFAULT false,
  "password" varchar NOT NULL,
  "location" varchar,
  "login_count" integer NOT NULL DEFAULT 1,
  "language" user_language NOT NULL DEFAULT 'en',
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "last_login" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "image" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "link" varchar NOT NULL,
  "thumbnail_link" varchar NOT NULL,
  "type" image_type NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "moniest" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar UNIQUE NOT NULL,
  "bio" varchar,
  "description" text,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "moniest_post_crypto_statistics" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "moniest_id" varchar NOT NULL,
  "pnl_7days" float,
  "roi_7days" float,
  "win_rate_7days" float,
  "posts_7days" varchar [],
  "pnl_30days" float,
  "roi_30days" float,
  "win_rate_30days" float,
  "posts_30days" varchar [],
  "pnl_total" float,
  "roi_total" float,
  "win_rate_total" float,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "moniest_payout_info" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "moniest_id" varchar NOT NULL,
  "source" payout_source NOT NULL DEFAULT 'BINANCE',
  "type" payout_type NOT NULL DEFAULT 'BINANCE_ID',
  "value" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "moniest_subscription_info" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "moniest_id" varchar UNIQUE NOT NULL,
  "fee" float NOT NULL DEFAULT 5,
  "message" varchar,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user_subscription" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "moniest_id" varchar NOT NULL,
  "active" boolean NOT NULL DEFAULT true,
  "latest_transaction_id" varchar,
  "subscription_start_date" timestamp NOT NULL DEFAULT (now()),
  "subscription_end_date" timestamp NOT NULL DEFAULT (now()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user_subscription_history" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "moniest_id" varchar NOT NULL,
  "transaction_id" varchar,
  "subscription_start_date" timestamp NOT NULL DEFAULT (now()),
  "subscription_end_date" timestamp NOT NULL DEFAULT (now()),
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "post_crypto" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "moniest_id" varchar NOT NULL,
  "market_type" post_crypto_market_type NOT NULL,
  "currency" varchar NOT NULL,
  "start_price" float NOT NULL,
  "duration" timestamp NOT NULL,
  "take_profit" float NOT NULL,
  "stop" float NOT NULL,
  "target1" float,
  "target2" float,
  "target3" float,
  "direction" direction NOT NULL,
  "leverage" int NOT NULL,
  "finished" boolean NOT NULL DEFAULT false,
  "status" post_crypto_status NOT NULL DEFAULT 'pending',
  "pnl" float NOT NULL,
  "roi" float NOT NULL,
  "last_operated_at" bigint NOT NULL,
  "hit_price" float,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "post_crypto_description" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "post_id" varchar UNIQUE NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "password_reset_token" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "token" varchar UNIQUE NOT NULL,
  "token_expiry" timestamp NOT NULL,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "email_verification_token" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar NOT NULL,
  "token" varchar UNIQUE NOT NULL,
  "token_expiry" timestamp NOT NULL,
  "redirect_url" varchar NOT NULL,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "feedback" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar,
  "type" varchar,
  "message" varchar NOT NULL,
  "solved" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "solved_at" timestamp
);

CREATE TABLE "binance_payment_transaction" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "qrcode_link" text NOT NULL,
  "checkout_link" text NOT NULL,
  "deep_link" text NOT NULL,
  "universal_link" text NOT NULL,
  "status" binance_payment_status NOT NULL DEFAULT 'pending',
  "user_id" varchar NOT NULL,
  "moniest_id" varchar NOT NULL,
  "date_type" binance_payment_date_type NOT NULL DEFAULT 'MONTH',
  "date_value" integer NOT NULL,
  "moniest_fee" float NOT NULL,
  "amount" float NOT NULL,
  "webhook_url" text NOT NULL,
  "payer_id" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "binance_payout_history" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "transaction_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "moniest_id" varchar NOT NULL,
  "payer_id" varchar NOT NULL,
  "total_amount" float NOT NULL,
  "amount" float NOT NULL,
  "date_type" binance_payment_date_type NOT NULL DEFAULT 'MONTH',
  "date_value" integer NOT NULL,
  "date_index" integer NOT NULL,
  "payout_date" timestamp NOT NULL,
  "payout_year" integer NOT NULL,
  "payout_month" integer NOT NULL,
  "payout_day" integer NOT NULL,
  "status" binance_payout_status NOT NULL DEFAULT 'pending',
  "operation_fee_percentage" float,
  "payout_done_at" timestamp,
  "payout_request_id" varchar,
  "failure_message" text,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "user" ("username");

CREATE UNIQUE INDEX ON "user" ("email");

CREATE UNIQUE INDEX ON "image" ("user_id", "type");

CREATE UNIQUE INDEX ON "moniest" ("user_id");

CREATE UNIQUE INDEX ON "moniest_post_crypto_statistics" ("moniest_id");

CREATE INDEX ON "moniest_post_crypto_statistics" ("pnl_7days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("roi_7days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("win_rate_7days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("posts_7days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("pnl_30days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("roi_30days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("win_rate_30days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("posts_30days");

CREATE INDEX ON "moniest_post_crypto_statistics" ("pnl_total");

CREATE INDEX ON "moniest_post_crypto_statistics" ("roi_total");

CREATE INDEX ON "moniest_post_crypto_statistics" ("win_rate_total");

CREATE UNIQUE INDEX ON "moniest_payout_info" ("moniest_id", "source");

CREATE INDEX ON "moniest_subscription_info" ("moniest_id");

CREATE UNIQUE INDEX ON "user_subscription" ("user_id", "moniest_id");

CREATE UNIQUE INDEX ON "user_subscription_history" ("user_id", "moniest_id", "transaction_id");

CREATE INDEX ON "post_crypto" ("moniest_id");

CREATE INDEX ON "post_crypto" ("finished");

CREATE INDEX ON "post_crypto" ("duration");

CREATE INDEX ON "post_crypto" ("created_at");

CREATE INDEX ON "post_crypto" ("last_operated_at");

CREATE INDEX ON "post_crypto_description" ("post_id");

CREATE UNIQUE INDEX ON "password_reset_token" ("user_id", "token");

CREATE UNIQUE INDEX ON "email_verification_token" ("user_id", "token");

CREATE UNIQUE INDEX ON "binance_payout_history" (
  "transaction_id",
  "user_id",
  "moniest_id",
  "payout_year",
  "payout_month",
  "payout_day"
);

COMMENT ON TABLE "user" IS 'Stores user data';

COMMENT ON TABLE "image" IS 'Stores image data';

COMMENT ON TABLE "moniest" IS 'Stores moniest data';

COMMENT ON TABLE "moniest_post_crypto_statistics" IS 'Stores moniest crypto statistics info';

COMMENT ON TABLE "moniest_payout_info" IS 'Stores moniest payout info';

COMMENT ON TABLE "moniest_subscription_info" IS 'Stores subscription data of a moniest';

COMMENT ON TABLE "user_subscription" IS 'Stores user subscription info';

COMMENT ON TABLE "user_subscription_history" IS 'Stores user subscriptions history';

COMMENT ON TABLE "post_crypto" IS 'Stores crypto posts data';

COMMENT ON TABLE "post_crypto_description" IS 'Stores crypto post description data';

COMMENT ON TABLE "password_reset_token" IS 'Stores reset token for forget password operations';

COMMENT ON TABLE "email_verification_token" IS 'Stores email verification token for verifying account';

COMMENT ON TABLE "feedback" IS 'Stores feedback from users';

COMMENT ON TABLE "binance_payment_transaction" IS 'Stores binance payment transactions info and history';

COMMENT ON TABLE "binance_payout_history" IS 'Stores binance payout info and history';

ALTER TABLE "image"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "moniest"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "moniest_subscription_info"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "moniest_post_crypto_statistics"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "moniest_payout_info"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "binance_payment_transaction"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "binance_payment_transaction"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "binance_payout_history"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "binance_payout_history"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "user_subscription"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_subscription"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "user_subscription_history"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_subscription_history"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "post_crypto"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "password_reset_token"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "email_verification_token"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "post_crypto_description"
ADD FOREIGN KEY ("post_id") REFERENCES "post_crypto" ("id");

ALTER TABLE "feedback"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");