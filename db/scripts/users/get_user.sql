--	Get(ctx context.Context, userID int64) (*User, error)
SELECT * FROM "users" WHERE "users".tg_id = 11111;

SELECT * FROM "statements" INNER JOIN 
    "users" ON "statements".user_id = "users".id
    WHERE "users".tg_id = 11111;

SELECT DISTINCT *
FROM "variables" INNER JOIN
    "statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
    "statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
    "users"                 ON "statements".user_id                 = "users".id
WHERE ("users".tg_id = 11111);
