--	Add(ctx context.Context, user *User) error
INSERT INTO users (tg_id, password, created_at)
VALUES (11111, 'password', now());
