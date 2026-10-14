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
  "entity_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "state_definitions_entity_kinds_state_definitions" FOREIGN KEY ("entity_id") REFERENCES "public"."entity_kinds" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "statedefinitions_key_entity_id" to table: "state_definitions"
CREATE UNIQUE INDEX "statedefinitions_key_entity_id" ON "public"."state_definitions" ("key", "entity_id");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "email" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "status_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_state_definitions_user" FOREIGN KEY ("status_id") REFERENCES "public"."state_definitions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
-- Create "tenants" table
CREATE TABLE "public"."tenants" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "slug" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "status_id" uuid NOT NULL,
  "created_by" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "tenants_state_definitions_tenants" FOREIGN KEY ("status_id") REFERENCES "public"."state_definitions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "tenants_users_tenant" FOREIGN KEY ("created_by") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "tenants_created_by_key" to table: "tenants"
CREATE UNIQUE INDEX "tenants_created_by_key" ON "public"."tenants" ("created_by");
-- Create index "tenants_slug_key" to table: "tenants"
CREATE UNIQUE INDEX "tenants_slug_key" ON "public"."tenants" ("slug");
-- Create "workspaces" table
CREATE TABLE "public"."workspaces" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "slug" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "status_id" uuid NOT NULL,
  "tenant_id" uuid NOT NULL,
  "created_by" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "workspaces_state_definitions_workspace" FOREIGN KEY ("status_id") REFERENCES "public"."state_definitions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "workspaces_tenants_workspace" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "workspaces_users_creator" FOREIGN KEY ("created_by") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "workspaces_tenant_id_key" to table: "workspaces"
CREATE UNIQUE INDEX "workspaces_tenant_id_key" ON "public"."workspaces" ("tenant_id");
-- Create "projects" table
CREATE TABLE "public"."projects" (
  "id" uuid NOT NULL,
  "name" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "status_id" uuid NOT NULL,
  "tenant_id" uuid NOT NULL,
  "workspace_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "projects_state_definitions_projects" FOREIGN KEY ("status_id") REFERENCES "public"."state_definitions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "projects_tenants_projects" FOREIGN KEY ("tenant_id") REFERENCES "public"."tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "projects_workspaces_projects" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "user_credentials" table
CREATE TABLE "public"."user_credentials" (
  "id" uuid NOT NULL,
  "password_hash" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_credentials_users_credentials" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "user_credentials_user_id_key" to table: "user_credentials"
CREATE UNIQUE INDEX "user_credentials_user_id_key" ON "public"."user_credentials" ("user_id");
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
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_sessions_users_sessions" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "user_sessions_refresh_token_hash_key" to table: "user_sessions"
CREATE UNIQUE INDEX "user_sessions_refresh_token_hash_key" ON "public"."user_sessions" ("refresh_token_hash");
-- Create index "usersession_user_id" to table: "user_sessions"
CREATE INDEX "usersession_user_id" ON "public"."user_sessions" ("user_id");
