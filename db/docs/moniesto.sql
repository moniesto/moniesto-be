CREATE TYPE "image_type" AS ENUM ('profile_photo', 'background_photo');

CREATE TYPE "entry_position" AS ENUM ('long', 'short');

CREATE TYPE "post_crypto_status" AS ENUM ('pending', 'fail', 'success');

CREATE TYPE "payout_source" AS ENUM ('BINANCE');

CREATE TYPE "payout_type" AS ENUM ('BINANCE_ID');

CREATE TABLE "user" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "email_verified" boolean NOT NULL DEFAULT false,
  "password" varchar NOT NULL,
  "location" varchar,
  "login_count" integer NOT NULL DEFAULT 1,
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
  "score" float NOT NULL DEFAULT 0,
  "deleted" boolean NOT NULL DEFAULT false,
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
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "post_crypto" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "moniest_id" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "start_price" float NOT NULL,
  "duration" timestamp NOT NULL,
  "target1" float NOT NULL,
  "target2" float NOT NULL,
  "target3" float NOT NULL,
  "stop" float NOT NULL,
  "direction" entry_position NOT NULL,
  "score" float NOT NULL DEFAULT 0,
  "finished" boolean NOT NULL DEFAULT false,
  "status" post_crypto_status NOT NULL DEFAULT 'pending',
  "last_target_hit" float NOT NULL,
  "last_job_timestamp" bigint NOT NULL,
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

CREATE TABLE "card" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "token" varchar,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user_card" (
  "user_id" varchar NOT NULL,
  "card_id" varchar UNIQUE NOT NULL
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

CREATE UNIQUE INDEX ON "user" ("username");

CREATE UNIQUE INDEX ON "user" ("email");

CREATE UNIQUE INDEX ON "image" ("user_id", "type");

CREATE UNIQUE INDEX ON "moniest" ("user_id");

CREATE UNIQUE INDEX ON "moniest_payout_info" ("moniest_id", "source");

CREATE INDEX ON "moniest_subscription_info" ("moniest_id");

CREATE UNIQUE INDEX ON "user_subscription" ("user_id", "moniest_id");

CREATE INDEX ON "post_crypto" ("moniest_id");

CREATE INDEX ON "post_crypto_description" ("post_id");

CREATE UNIQUE INDEX ON "user_card" ("user_id", "card_id");

CREATE UNIQUE INDEX ON "password_reset_token" ("user_id", "token");

CREATE UNIQUE INDEX ON "email_verification_token" ("user_id", "token");

COMMENT ON TABLE "user" IS 'Stores user data';

COMMENT ON TABLE "image" IS 'Stores image data';

COMMENT ON TABLE "moniest" IS 'Stores moniest data';

COMMENT ON TABLE "moniest_payout_info" IS 'Stores moniest payout info';

COMMENT ON TABLE "moniest_subscription_info" IS 'Stores subscription data of a moniest';

COMMENT ON TABLE "user_subscription" IS 'Stores user subscription info';

COMMENT ON TABLE "post_crypto" IS 'Stores crypto posts data';

COMMENT ON TABLE "post_crypto_description" IS 'Stores crypto post description data';

COMMENT ON TABLE "card" IS 'Stores single card data';

COMMENT ON TABLE "user_card" IS 'Stores relation between user and card';

COMMENT ON TABLE "password_reset_token" IS 'Stores reset token for forget password operations';

COMMENT ON TABLE "email_verification_token" IS 'Stores email verification token for verifying account';

COMMENT ON TABLE "feedback" IS 'Stores feedback from users';

ALTER TABLE "image"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "moniest"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_card"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_card"
ADD FOREIGN KEY ("card_id") REFERENCES "card" ("id");

ALTER TABLE "moniest_subscription_info"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "moniest_payout_info"
ADD FOREIGN KEY ("moniest_id") REFERENCES "moniest" ("id");

ALTER TABLE "user_subscription"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_subscription"
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