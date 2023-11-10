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

CREATE TABLE IF NOT EXISTS "consultations" (
  "id" UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  "doctorId" UUID NOT NULL REFERENCES doctors(id),
  "patientId" UUID NOT NULL REFERENCES patients(id),
  "date" TIMESTAMP NOT NULL,
  "description" VARCHAR(250),
  "notes" VARCHAR(500),
  "diagnosis" VARCHAR(1024),
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
)

