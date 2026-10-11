CREATE TABLE "users" (
  "id"         uuid              NOT NULL,
  "email"      character varying NOT NULL,
  "created_at" timestamptz       NOT NULL,
  "updated_at" timestamptz       NOT NULL,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");

CREATE TABLE "user_credentials" (
  "id"               uuid              NOT NULL,
  "password_hash"    character varying NOT NULL,
  "user_credentials" uuid              NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_credentials_users_credentials"
    FOREIGN KEY ("user_credentials") REFERENCES "users" ("id")
    ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- unique FK column enforces the one-to-one edge
CREATE UNIQUE INDEX "user_credentials_user_credentials_key"
  ON "user_credentials" ("user_credentials");
  