-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "username" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS schedules_id_seq;

-- Table Definition
CREATE TABLE "public"."schedules" (
    "id" int8 NOT NULL DEFAULT nextval('schedules_id_seq'::regclass),
    "cinema_id" int8,
    "movie_id" int8,
    "show_date" timestamptz,
    "start_time" text,
    "end_time" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);