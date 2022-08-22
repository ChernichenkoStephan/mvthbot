--	DeleteWithNames(ctx context.Context, userID int64, names []string) error
UPDATE variables
SET value = NULL
WHERE variables.id IN (
    SELECT variables.id
    FROM variables INNER JOIN
        statementsVariables   ON statementsVariables.variable_id    = variables.id INNER JOIN
        statements            ON statementsVariables.statement_id   = statements.id INNER JOIN
        users                 ON statements.user_id                 = users.id

    WHERE users.tg_id = 11111 AND (variables.name IN ('b', 'c'))
) RETURNING id;