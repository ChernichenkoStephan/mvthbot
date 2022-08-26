ALTER TABLE "statements" DROP CONSTRAINT "FK_user_id";
ALTER TABLE "statementsVariables" DROP CONSTRAINT "FK_variable_id";
ALTER TABLE "statementsVariables" DROP CONSTRAINT "FK_statement_id";


TRUNCATE TABLE "variables";
TRUNCATE TABLE "statements";
TRUNCATE TABLE "statementsVariables" CASCADE;
TRUNCATE TABLE "users" CASCADE;


ALTER TABLE "statements" ADD CONSTRAINT "FK_user_id"
FOREIGN KEY("user_id") REFERENCES "users"("id");

ALTER TABLE "statementsVariables" ADD CONSTRAINT "FK_variable_id"
FOREIGN KEY("variable_id") REFERENCES "variables"("id");

ALTER TABLE "statementsVariables" ADD CONSTRAINT "FK_statement_id"
FOREIGN KEY("statement_id") REFERENCES "statements"("id");


