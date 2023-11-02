CREATE DATABASE "medicalapi";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "patients" (
  "id" UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL UNIQUE,
  "document" VARCHAR(11) NOT NULL,
  "birth_date" DATE NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "doctors" (
	"id" UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
	"name" VARCHAR(255) NOT NULL,
	"email" VARCHAR(255) NOT NULL UNIQUE,
	"password" VARCHAR(255) NOT NULL,
	"document" VARCHAR(11) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMP NOT NULL DEFAULT now()
)

