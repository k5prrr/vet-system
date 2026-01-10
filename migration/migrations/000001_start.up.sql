
-- clients
DROP SEQUENCE IF EXISTS clients_id_seq;
CREATE SEQUENCE clients_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."clients" (
    "id" bigint DEFAULT nextval('clients_id_seq') NOT NULL,
    "fio" varchar NOT NULL,
    "phone" varchar(255) NOT NULL UNIQUE,
    "birth_date" timestamptz,
    "parent_id" bigint NOT NULL,
    "description" text NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    CONSTRAINT "clients_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);


-- timesheet
DROP SEQUENCE IF EXISTS timesheet_id_seq;
CREATE SEQUENCE timesheet_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."timesheet" (
    "id" bigint DEFAULT nextval('timesheet_id_seq') NOT NULL,
    "user_id" bigint NOT NULL,
    "parent_id" bigint NOT NULL,
    "date" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    CONSTRAINT "timesheet_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);



-- Animals animal_types
DROP SEQUENCE IF EXISTS animals_id_seq;
CREATE SEQUENCE animals_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;
CREATE TABLE "public"."animals" (
    "id" bigint DEFAULT nextval('animals_id_seq') NOT NULL,
    "name" varchar NOT NULL,
    "birth_date" date,
    "chip" varchar,
    "client_id" bigint,
    "animal_type_id" bigint NOT NULL,
    "description" text,
    "parent_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "animals_pkey" PRIMARY KEY ("id")
)
    WITH (oids = false);
CREATE UNIQUE INDEX animals_chip ON public.animals USING btree (chip);
CREATE INDEX animals_client_id ON public.animals USING btree (client_id);
CREATE INDEX animals_name ON public.animals USING btree (name);

DROP SEQUENCE IF EXISTS animal_types_id_seq;
CREATE SEQUENCE animal_types_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."animal_types" (
    "id" bigint DEFAULT nextval('animal_types_id_seq') NOT NULL,
    "name" varchar NOT NULL,
    "code" varchar NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "animal_types_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);
INSERT INTO "animal_types" ("id", "name", "code", "created_at", "updated_at") VALUES
(1,	'кошка',	'cat', NOW(), NOW()),
(2,	'собака',	'dog', NOW(), NOW()),
(3,	'пернатые',	'bird', NOW(), NOW()),
(4,	'иные',	'other', NOW(), NOW());








-- records record_status
DROP SEQUENCE IF EXISTS records_id_seq;
CREATE SEQUENCE records_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."records" (
    "id" bigint DEFAULT nextval('records_id_seq') NOT NULL,
    "timesheet_id" bigint NOT NULL,
    "date_time" timestamptz NOT NULL,
    "client_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    "parent_id" bigint NOT NULL,
    "parent_role_id" bigint NOT NULL,
    "status_id" bigint NOT NULL,
    "animal_id" bigint NOT NULL,
    "complaints" text,
    "examination" text NOT NULL,
    "ds" text,
    "recommendations" text NOT NULL,
    "description" text,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "records_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

DROP SEQUENCE IF EXISTS record_status_id_seq;
CREATE SEQUENCE record_status_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."record_status" (
    "id" bigint DEFAULT nextval('record_status_id_seq') NOT NULL,
    "name" varchar NOT NULL,
    "code" varchar NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamp,
    CONSTRAINT "record_status_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

INSERT INTO "record_status" ("id", "name", "code", "created_at", "updated_at") VALUES
    (1,	'новый',	'new', NOW(), NOW()),
    (2,	'подтвержден',	'approve', NOW(), NOW()),
    (3,	'ожидает приема',	'wait', NOW(), NOW()),
    (4,	'завершен',	'complited', NOW(), NOW()),
    (5,	'отменён',	'cancel', NOW(), NOW()),
    (6,	'не явился',	'non', NOW(), NOW());














-- users user_roles
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."users" (
    "id" bigint DEFAULT nextval('users_id_seq') NOT NULL,
    "fio" varchar NOT NULL,
    "role_id" bigint NOT NULL,
    "phone" varchar NOT NULL,
    "parent_id" bigint NOT NULL,
    "description" text,
    "password_hash" varchar NOT NULL,
    "auth_secret" varchar NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

DROP SEQUENCE IF EXISTS user_roles_id_seq;
CREATE SEQUENCE user_roles_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."user_roles" (
   "id" bigint DEFAULT nextval('user_roles_id_seq') NOT NULL,
   "name" varchar NOT NULL,
   "code" varchar NOT NULL,
   "created_at" timestamptz NOT NULL,
   "updated_at" timestamptz NOT NULL,
   "deleted_at" timestamp,
   CONSTRAINT "user_roles_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);

INSERT INTO "user_roles" ("id", "name", "code", "created_at", "updated_at") VALUES
    (1,	'без роли',	'non', NOW(), NOW()),
    (2,	'Клиент',	'client', NOW(), NOW()),
    (3,	'Доктор',	'doctor', NOW(), NOW()),
    (4,	'Админ',	'admin', NOW(), NOW());


CREATE UNIQUE INDEX users_phone ON public.users USING btree (phone);

CREATE INDEX users_fio ON public.users USING btree (fio);

CREATE INDEX unique_client_phone ON public.clients USING btree (phone);

CREATE INDEX users_role_id ON public.users USING btree (role_id);

-- Для поиска записей врача в день/неделю
CREATE INDEX records_user_id_date_time ON records (user_id, date_time);

-- Для поиска записей клиента (личный кабинет)
CREATE INDEX records_client_id ON records (client_id, date_time DESC);

-- Для поиска активных (неудалённых, неотменённых) записей
CREATE INDEX records_active ON records (status_id) WHERE deleted_at IS NULL AND status_id NOT IN (5,6);


