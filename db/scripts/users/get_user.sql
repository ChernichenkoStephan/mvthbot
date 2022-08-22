--	Get(ctx context.Context, userID int64) (*User, error)
SELECT * FROM users WHERE users.tg_id = 11111;

SELECT  statements.id, statements.equation, statements.value, statements.created_at
FROM statements INNER JOIN 
    users ON statements.user_id = users.id
    WHERE users.tg_id = 11111;

SELECT DISTINCT variables.name, variables.value
FROM variables INNER JOIN
    "statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
    statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
    users                 ON statements.user_id                 = users.id
WHERE (users.tg_id = 11111);
