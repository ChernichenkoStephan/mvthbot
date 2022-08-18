--	GetWithNames(ctx context.Context, userID int64, names []string) (VMap, error)
SELECT DISTINCT variables.name, variables.value 
FROM "variables" INNER JOIN
    "statementsVariables"   ON "statementsVariables".variable_id    = "variables".id INNER JOIN
    "statements"            ON "statementsVariables".statement_id   = "statements".id INNER JOIN
    "users"                 ON "statements".user_id                 = "users".id
WHERE "variables".name IN ('a', 'b') AND "users".tg_id = 11111;
