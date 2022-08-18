ALTER TABLE "statements" DROP CONSTRAINT "FK_user_id";
ALTER TABLE "statementsVariables" DROP CONSTRAINT "FK_variable_id";
ALTER TABLE "statementsVariables" DROP CONSTRAINT "FK_statement_id";

DROP TABLE IF EXISTS "statements";
DROP TABLE IF EXISTS "variables";
DROP TABLE IF EXISTS "statementsVariables";
DROP TABLE IF EXISTS "users";
