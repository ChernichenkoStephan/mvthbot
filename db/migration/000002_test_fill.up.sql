-- id=1
INSERT INTO users (tg_id, password, created_at)
VALUES (11111, 'password', now());


-- id=1 a=2+2
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, '2+2', 4.0, now());

-- id=2 b=1+
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, '1+1', 2.0, now());

-- id=3 b=c=2*2
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (1, '2*2', 4.0, now());


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

-- //////////////////////////////////////////////////////////

-- id=2
INSERT INTO users (tg_id, password, created_at)
VALUES (22222, '22222', now());


-- id=4 a=7+7
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (2, '7+7', 14.0, now());

-- id=5 b=10+10
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (2, '10+10', 20.0, now());

-- id=6 b=c=20*2
INSERT INTO statements (user_id, equation, value, created_at)
VALUES (2, '20*2', 40.0, now());


-- id=4
INSERT INTO variables (name, value, created_at)
values ('a', 14.0, now());

-- id=5
insert into variables (name, value, created_at)
values ('b', 40.0, now());

-- id=6
INSERT INTO variables (name, value, created_at)
values ('c', 40.0, now());

INSERT INTO "statementsVariables" (variable_id, statement_id)
values (4, 4);

INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (5, 5);

INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (5, 6);

INSERT INTO "statementsVariables" (variable_id, statement_id)
VALUES (6, 6);
