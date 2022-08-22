INSERT INTO statements (user_id, equation, value, created_at)
	VALUES ((SELECT id FROM users WHERE users.tg_id = 11111), 'a=c=d=2*2*2', 8.0, now()) RETURNING id;

SELECT * FROM set_var(11111, 7, 'a', 8.0)
    UNION SELECT * FROM set_var(11111, 7, 'c', 8.0)
    UNION SELECT * FROM set_var(11111, 7, 'd', 8.0);
