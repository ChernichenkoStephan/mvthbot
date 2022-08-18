CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "tg_id" numeric NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "statements" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "equation" text NOT NULL,
  "value" double precision NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variables" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "value" double precision,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "statementsVariables" (
  "variable_id" bigserial NOT NULL,
  "statement_id" bigserial NOT NULL,
  PRIMARY KEY("variable_id", "statement_id")
);


ALTER TABLE "statements" ADD CONSTRAINT "FK_user_id"
FOREIGN KEY("user_id") REFERENCES "users"("id");

ALTER TABLE "statementsVariables" ADD CONSTRAINT "FK_variable_id"
FOREIGN KEY("variable_id") REFERENCES "variables"("id") ON DELETE CASCADE;

ALTER TABLE "statementsVariables" ADD CONSTRAINT "FK_statement_id"
FOREIGN KEY("statement_id") REFERENCES "statements"("id") ON DELETE CASCADE;


CREATE INDEX ON "variables" ("name");

CREATE INDEX ON "statementsVariables" ("variable_id");

CREATE INDEX ON "statementsVariables" ("statement_id");

CREATE INDEX ON "statementsVariables" ("variable_id", "statement_id");

COMMENT ON COLUMN "users"."id" IS 'same as telegram user id';

COMMENT ON COLUMN "variables"."value" IS 'if value equals null, then variable is deleted. (drop line only if statement dropped)';

GRANT ALL PRIVILEGES ON TABLE "users" TO admin;
GRANT ALL PRIVILEGES ON TABLE "statements" TO admin;
GRANT ALL PRIVILEGES ON TABLE "variables" TO admin;
GRANT ALL PRIVILEGES ON TABLE "statementsVariables" TO admin;


GRANT USAGE, SELECT ON SEQUENCE statements_id_seq TO admin;
GRANT USAGE, SELECT ON SEQUENCE variables_id_seq TO admin;
GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO admin;
