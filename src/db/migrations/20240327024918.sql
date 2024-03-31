-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
