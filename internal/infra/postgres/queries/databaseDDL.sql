CREATE DATABASE "medicalapi";

CREATE TABLE IF NOT EXISTS "patients" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "document" VARCHAR(11) NOT NULL,
    "birth_date" DATE NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

