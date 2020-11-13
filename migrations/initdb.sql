-- DROP TABLE IF EXISTS "public"."users";
-- -- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.
--
-- CREATE SEQUENCE IF NOT EXISTS id_seq;
--
-- CREATE TABLE "public"."users" (
--     "id" SERIAL,
--     "username" varchar(256),
--     "email" varchar(256),
--     "password" varchar(64),
--     "created_at" timestamp NOT NULL DEFAULT '-infinity'::timestamp without time zone,
--     "updated_at" timestamp NOT NULL DEFAULT '-infinity'::timestamp without time zone,
--     PRIMARY KEY ("id")
-- );
--
-- COMMENT ON COLUMN "public"."users"."id" IS 'id of user';
-- COMMENT ON COLUMN "public"."users"."username" IS 'user name';
-- COMMENT ON COLUMN "public"."users"."email" IS 'email of user';
-- COMMENT ON COLUMN "public"."users"."password" IS 'password of user';
--
-- INSERT INTO "public"."users" ("id", "username", "email", "password", "created_at", "updated_at")
-- VALUES
-- ('1', 'nhinhdt', 'nhinhdt@tmh-techlab.vn', '12345', '2015-05-15 13:51:07.854938', '2015-05-15 13:51:07.854938'),
-- ('2', 'tranglt', 'j-nagano@tribalmedia.co.jp', '12423345456', '2015-05-15 13:51:07.854938', '2015-05-15 13:51:07.854938'),
-- ('3', 't-segawa+em', 't-segawa+em_mv@tribalmedia.co.jp', '7c4a8d09ca3762af61e59520943dc26494f8941b', '2015-05-15 13:51:07.854938', '2015-05-15 13:51:07.854938');




