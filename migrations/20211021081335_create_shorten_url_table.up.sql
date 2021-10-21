CREATE TABLE "shorten_url" (
  "id" serial PRIMARY KEY,
  "short_url" text UNIQUE NOT NULL,
  "url" text UNIQUE NOT NULL,
  "created_at" timestamp with time zone NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp with time zone NOT NULL DEFAULT (CURRENT_TIMESTAMP)
)