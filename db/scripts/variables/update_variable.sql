--	Update(ctx context.Context, userID int64, name string, value float64) error
UPDATE variables
    SET value = 321.0
    WHERE id = (
        SELECT  variables.id
        FROM variables INNER JOIN
            "statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
            statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
            users                 ON statements.user_id                 = users.id
    WHERE users.tg_id = 11111 AND variables.name = 'b'
    ) RETURNING id;
