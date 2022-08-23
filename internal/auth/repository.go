package auth

import (
	"context"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	cache user.Cache
	db    *sqlx.DB
}

func NewAuthRepository(cache user.Cache, db *sqlx.DB) *authRepository {
	return &authRepository{
		cache: cache,
		db:    db,
	}
}

type dbUser struct {
	Id       int    `db:"id"`
	TgID     int64  `db:"tg_id"`
	Password string `db:"password"`
	Created  string `db:"created_at"`
}

func (repo *authRepository) GetPassword(ctx context.Context, tgID int64) (string, error) {

	query := `SELECT * FROM "users" WHERE "users".tg_id = $1;`

	u := dbUser{}
	err := repo.db.Get(&u, query, tgID)
	if err != nil {
		return "", errors.Wrap(err, "Get request fail")
	}

	return u.Password, nil
}
