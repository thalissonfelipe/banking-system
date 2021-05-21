BEGIN;

CREATE TABLE IF NOT EXISTS "account" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "cpf" VARCHAR(50) NOT NULL,
    "secret" VARCHAR(200) NOT NULL,
    "balance" INT NOT NULL,
    "created_at" DATE
);

COMMIT;