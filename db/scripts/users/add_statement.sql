--	AddStatement(ctx context.Context, userID int64, statement *solv.Statement) error

INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, 'd=e=3*3', 9.0, now());

INSERT INTO variables (name, value, created_at)
values ('d', 9.0, now());

INSERT INTO variables (name, value, created_at)
values ('e', 9.0, now());


INSERT INTO "statementsVariables" (variable_id, statement_id)
values (4, 4);

INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (5, 4);


select * from "users";
select * from "statements";
select * from "variables";
select * from "statementsVariables";
