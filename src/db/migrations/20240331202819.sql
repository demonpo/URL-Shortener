-- Create "shorteners" table
CREATE TABLE "public"."shorteners" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "url" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_shorteners_deleted_at" to table: "shorteners"
CREATE INDEX "idx_shorteners_deleted_at" ON "public"."shorteners" ("deleted_at");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "uni_users_email" to table: "users"
CREATE UNIQUE INDEX "uni_users_email" ON "public"."users" ("email");
