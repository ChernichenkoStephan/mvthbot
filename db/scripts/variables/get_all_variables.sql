--	GetAll(ctx context.Context, userID int64) (VMap, error)
SELECT DISTINCT variables.name as "name", variables.value as "value"
FROM "variables" INNER JOIN
    "statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
    statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
    users                 ON statements.user_id                 = users.id
WHERE users.tg_id = 22222 AND variables.value IS NOT NULL;
