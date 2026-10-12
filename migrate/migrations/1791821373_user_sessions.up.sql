-- Create "user_sessions" table
CREATE TABLE "user_sessions" (
  "id" uuid NOT NULL,
  "refresh_token_hash" character varying NOT NULL,
  "user_agent" character varying(512) NULL,
  "expires_at" timestamptz NOT NULL,
  "last_used_at" timestamptz NULL,
  "revoked_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_sessions_users_sessions" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "user_sessions_refresh_token_hash_key" to table: "user_sessions"
CREATE UNIQUE INDEX "user_sessions_refresh_token_hash_key" ON "user_sessions" ("refresh_token_hash");
-- Create index "usersession_user_id" to table: "user_sessions"
CREATE INDEX "usersession_user_id" ON "user_sessions" ("user_id");