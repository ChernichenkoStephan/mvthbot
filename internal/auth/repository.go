package auth

import (
	"context"
	"fmt"
	"sync"

	"github.com/ChernichenkoStephan/mvthbot/internal/user"
)

type userAuthRepository struct {
	mx sync.Mutex
	db *user.IMStorage
}

func NewUserAuthRepository() *userAuthRepository {
	db := user.GetDefaultStorage()
	return &userAuthRepository{
		db: db,
	}
}

func (repo *userAuthRepository) GetPassword(ctx context.Context, id int64) (string, error) {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	pass, ok := repo.db.GetAuth(id)
	if !ok {
		return "", fmt.Errorf("UserNotFound")
	}
	return pass, nil
}
