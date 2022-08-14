package auth

import (
	"context"
	"fmt"

	"github.com/ChernichenkoStephan/mvthbot/internal/user"
)

type authRepository struct {
	cache *user.Cache
}

func NewAuthRepository(c *user.Cache) *authRepository {
	return &authRepository{
		cache: c,
		//db: db,
	}
}

func (repo *authRepository) GetPassword(ctx context.Context, id int64) (string, error) {
	it, ok := repo.cache.Get(fmt.Sprintf("%v", id))
	if !ok {
		return "", fmt.Errorf("UserNotFound")
	}

	u, ok := it.(*user.User)
	return u.Password, nil
}
