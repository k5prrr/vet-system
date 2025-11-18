-- Adminer 5.4.1 PostgreSQL 15.2 dump

DROP TABLE IF EXISTS "animal_tupes";
DROP SEQUENCE IF EXISTS animal_tupes_id_seq;
CREATE SEQUENCE animal_tupes_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."animal_tupes" (
    "id" integer DEFAULT nextval('animal_tupes_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    "code" character varying NOT NULL,
    CONSTRAINT "animal_tupes_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

INSERT INTO "animal_tupes" ("id", "name", "code") VALUES
(1,	'кошка',	'cat'),
(2,	'собака',	'dog'),
(3,	'пернатые',	'bird'),
(4,	'иные',	'other');

DROP TABLE IF EXISTS "animals";
DROP SEQUENCE IF EXISTS animals_id_seq;
CREATE SEQUENCE animals_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."animals" (
    "id" integer DEFAULT nextval('animals_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    "birth_date" date,
    "chip" character varying,
    "client_id" integer,
    "animal_tupe_id" integer NOT NULL,
    "description" text,
    "parent_id" integer NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "animals_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

CREATE UNIQUE INDEX animals_chip ON public.animals USING btree (chip);

CREATE INDEX animals_client_id ON public.animals USING btree (client_id);

CREATE INDEX animals_name ON public.animals USING btree (name);


DROP TABLE IF EXISTS "clients";
DROP SEQUENCE IF EXISTS clients_id_seq;
CREATE SEQUENCE clients_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."clients" (
    "id" integer DEFAULT nextval('clients_id_seq') NOT NULL,
    "fio" character varying NOT NULL,
    "phone" character(255) NOT NULL,
    "birth_date" date,
    "parent_id" integer NOT NULL,
    "description" text NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    CONSTRAINT "clients_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


DROP TABLE IF EXISTS "records";
DROP SEQUENCE IF EXISTS records_id_seq;
CREATE SEQUENCE records_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."records" (
    "id" integer DEFAULT nextval('records_id_seq') NOT NULL,
    "timesheet_id" integer NOT NULL,
    "date_time" timestamptz NOT NULL,
    "client_id" integer NOT NULL,
    "user_id" integer NOT NULL,
    "parent_id" integer NOT NULL,
    "parent_tupe" integer NOT NULL,
    "status_id" integer NOT NULL,
    "animal_id" integer NOT NULL,
    "complaints" text,
    "examination" text NOT NULL,
    "ds" text,
    "recommendations" text NOT NULL,
    "description" text,
    CONSTRAINT "records_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


DROP TABLE IF EXISTS "roles";
DROP SEQUENCE IF EXISTS roles_id_seq;
CREATE SEQUENCE roles_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."roles" (
    "id" integer DEFAULT nextval('roles_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    "code" character varying NOT NULL,
    CONSTRAINT "roles_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

INSERT INTO "roles" ("id", "name", "code") VALUES
(1,	'без роли',	'non'),
(2,	'Клиент',	'client'),
(3,	'Доктор',	'doctor'),
(4,	'Админ',	'admin');

DROP TABLE IF EXISTS "status_id";
DROP SEQUENCE IF EXISTS status_id_id_seq;
CREATE SEQUENCE status_id_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."status_id" (
    "id" integer DEFAULT nextval('status_id_id_seq') NOT NULL,
    "name" character varying NOT NULL,
    "code" character varying NOT NULL,
    CONSTRAINT "status_id_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

INSERT INTO "status_id" ("id", "name", "code") VALUES
(1,	'новый',	'new'),
(2,	'подтвержден',	'approve'),
(3,	'ожидает приема',	'wait'),
(4,	'завершен',	'complited'),
(5,	'отменён',	'cancel'),
(6,	'не явился',	'non');

DROP TABLE IF EXISTS "timesheet";
DROP SEQUENCE IF EXISTS timesheet_id_seq;
CREATE SEQUENCE timesheet_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."timesheet" (
    "id" integer DEFAULT nextval('timesheet_id_seq') NOT NULL,
    "user_id" integer NOT NULL,
    "parent_id" integer NOT NULL,
    "date" date NOT NULL,
    "created_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    CONSTRAINT "timesheet_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."users" (
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "fio" character varying NOT NULL,
    "role_id" integer NOT NULL,
    "phone" character varying NOT NULL,
    "parent_id" integer NOT NULL,
    "description" text,
    "token" character(255) NOT NULL,
    "password_hash" character(255) NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

CREATE UNIQUE INDEX users_phone ON public.users USING btree (phone);

CREATE INDEX users_fio ON public.users USING btree (fio);

CREATE INDEX users_role_id ON public.users USING btree (role_id);


-- 2025-11-18 16:36:31 UTC
