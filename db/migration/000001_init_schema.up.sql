CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "tg_id" numeric NOT NULL UNIQUE,
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

CREATE OR REPLACE FUNCTION set_var(u_tg_id BIGINT, statement_id BIGINT, var_name VARCHAR, var_val NUMERIC) RETURNS BIGINT AS $BODY$

DECLARE var_id integer;

BEGIN

    IF (SELECT COUNT(*)
        FROM "variables" INNER JOIN
            "statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
            "statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
            "users"                 ON "statements".user_id                 = "users".id
        WHERE "variables".name = var_name AND "users".tg_id = u_tg_id) THEN

        SELECT "variables".id FROM "variables" INNER JOIN
                    "statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
                    "statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
                    "users"                 ON "statements".user_id                 = "users".id
            WHERE "users".tg_id = u_tg_id AND "variables".name = var_name INTO var_id;

        UPDATE "variables"
            SET value = var_val
            WHERE id = var_id;

        RETURN var_id;

    ELSE 
        WITH inserted_id AS (
            INSERT INTO variables (name, value, created_at)
            VALUES (var_name, var_val, now()) RETURNING id
        ) SELECT * FROM inserted_id INTO var_id;

        INSERT INTO "statementsVariables" (variable_id, statement_id)
        VALUES (var_id, statement_id);

       RETURN var_id;

    END IF;

END
$BODY$ LANGUAGE 'plpgsql';
