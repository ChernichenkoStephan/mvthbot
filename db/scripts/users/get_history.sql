--	GetHistory(ctx context.Context, userID int64) (*History, error)
SELECT * FROM "statements" INNER JOIN 
    "users" ON "statements".user_id = "users".id
    WHERE "users".tg_id = 11111;
