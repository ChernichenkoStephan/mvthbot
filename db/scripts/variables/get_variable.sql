--	Get(ctx context.Context, userID int64, name string) (float64, error)
SELECT DISTINCT variables.name as "name", variables.value as "value"
FROM "variables" INNER JOIN
    "statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
    "statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
    "users"                 ON "statements".user_id                 = "users".id
WHERE "variables".name = 'a' AND "users".tg_id = 11111;
