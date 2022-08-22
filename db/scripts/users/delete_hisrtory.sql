--	DeleteHistory(ctx context.Context, userID int64) error

-- clearing variables
DELETE FROM variables 
WHERE variables.id IN (
    SELECT variables.id
    FROM variables INNER JOIN
        "statementsVariables"   ON "statementsVariables".variable_id    = variables.id INNER JOIN
        statements            ON "statementsVariables".statement_id   = statements.id INNER JOIN
        users                 ON statements.user_id                 = users.id
    WHERE users.tg_id = 11111
);

-- clearing statements
DELETE FROM statements 
WHERE statements.id IN (
    SELECT statements.id
    FROM statements INNER JOIN users ON statements.user_id = users.id
    WHERE users.tg_id = 11111
);
