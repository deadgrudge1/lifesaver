CREATE TABLE IF NOT EXISTS "users"(
    "id" varchar(64) primary key,
    "email" varchar(64),
    "name" varchar(64)
);

CREATE INDEX ON "users" ("email");