--	Update(ctx context.Context, user *User) error
UPDATE "users"
SET password = 'newpass'
WHERE "users".tg_id = 11111;
