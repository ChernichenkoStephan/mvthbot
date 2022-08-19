package auth

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	cache *user.Cache
	db    *sqlx.DB
}

func NewAuthRepository(c *user.Cache, db *sqlx.DB) *authRepository {
	return &authRepository{
		cache: c,
		db:    db,
	}
}

func (repo *authRepository) GetPassword(ctx context.Context, id int64) (string, error) {
	it, err := repo.cache.Get(fmt.Sprintf("%v", id))
	if !errors.As(err, &user.ItemNotFoundError{}) {
		return "", errors.Wrap(err, fmt.Sprintf("Getting user with id %d fail", id))
	} else if err == nil {
		u, ok := it.(*user.User)
		if !ok {
			return "", errors.New("Wrong type")
		}
		return u.Password, nil
	}
	return "", nil
}
