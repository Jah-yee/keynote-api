-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "email" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
-- Create "user_credentials" table
CREATE TABLE "public"."user_credentials" (
  "id" uuid NOT NULL,
  "password_hash" character varying NOT NULL,
  "user_credentials" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_credentials_users_credentials" FOREIGN KEY ("user_credentials") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "user_credentials_user_credentials_key" to table: "user_credentials"
CREATE UNIQUE INDEX "user_credentials_user_credentials_key" ON "public"."user_credentials" ("user_credentials");
-- Create "user_sessions" table
CREATE TABLE "public"."user_sessions" (
  "id" uuid NOT NULL,
  "refresh_token_hash" character varying NOT NULL,
  "user_agent" character varying NULL,
  "expires_at" timestamptz NOT NULL,
  "last_used_at" timestamptz NULL,
  "revoked_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "user_sessions" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_sessions_users_sessions" FOREIGN KEY ("user_sessions") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "user_sessions_refresh_token_hash_key" to table: "user_sessions"
CREATE UNIQUE INDEX "user_sessions_refresh_token_hash_key" ON "public"."user_sessions" ("refresh_token_hash");
-- Create index "usersession_user_sessions" to table: "user_sessions"
CREATE INDEX "usersession_user_sessions" ON "public"."user_sessions" ("user_sessions");
