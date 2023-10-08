-- Create "todos" table
CREATE TABLE "todos" ("id" uuid NOT NULL, "title" text NOT NULL, "description" text NOT NULL, "image" text NOT NULL, "starts_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "ends_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "tags" table
CREATE TABLE "tags" ("id" uuid NOT NULL, "keyword" text NOT NULL, "todo_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "tags_todos_tags" FOREIGN KEY ("todo_id") REFERENCES "todos" ("id") ON DELETE NO ACTION);
