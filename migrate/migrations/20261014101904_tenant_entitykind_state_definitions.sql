-- Modify "user_credentials" table
ALTER TABLE "public"."user_credentials" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "deleted_at" timestamptz NULL;
-- Create "entity_kinds" table
CREATE TABLE "public"."entity_kinds" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "key" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "entity_kinds_key_key" to table: "entity_kinds"
CREATE UNIQUE INDEX "entity_kinds_key_key" ON "public"."entity_kinds" ("key");
-- Create "state_definitions" table
CREATE TABLE "public"."state_definitions" (
  "id" uuid NOT NULL,
  "key" character varying NOT NULL,
  "name" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "entity_kind_state_definitions" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "state_definitions_entity_kinds_state_definitions" FOREIGN KEY ("entity_kind_state_definitions") REFERENCES "public"."entity_kinds" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "statedefinitions_key_entity_kind_state_definitions" to table: "state_definitions"
CREATE UNIQUE INDEX "statedefinitions_key_entity_kind_state_definitions" ON "public"."state_definitions" ("key", "entity_kind_state_definitions");
-- Create "tenants" table
CREATE TABLE "public"."tenants" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "slug" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "state_definitions_tenants" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "tenants_state_definitions_tenants" FOREIGN KEY ("state_definitions_tenants") REFERENCES "public"."state_definitions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "tenants_slug_key" to table: "tenants"
CREATE UNIQUE INDEX "tenants_slug_key" ON "public"."tenants" ("slug");
