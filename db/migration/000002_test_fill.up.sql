-- id=1
INSERT INTO users (tg_id, password, created_at)
VALUES (11111, 'password', now());


-- id=1
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, 'a=2+2', 4.0, now());

-- id=2
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, 'b=1+1', 2.0, now());

-- id=3
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, 'b=c=2*2', 4.0, now());


-- id=1
INSERT INTO variables (name, value, created_at)
values ('a', 4.0, now());

-- id=2
insert into variables (name, value, created_at)
values ('b', 4.0, now());

-- id=3
INSERT INTO variables (name, value, created_at)
values ('c', 4.0, now());


-- a=2+2
INSERT INTO "statementsVariables" (variable_id, statement_id)
values (1, 1);

-- b=1+1
INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (2, 2);

-- b=c=2*2
INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (2, 3);

-- b=c=2*2
INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (3, 3);

