-- Create "clicks" table
CREATE TABLE "public"."clicks" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "shortener_id" text NULL,
  "user_ip" text NULL,
  "referrer_url" text NULL,
  "user_agent" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_clicks_deleted_at" to table: "clicks"
CREATE INDEX "idx_clicks_deleted_at" ON "public"."clicks" ("deleted_at");
