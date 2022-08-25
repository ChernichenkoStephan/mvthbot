package auth

import (
	"context"

	"go.uber.org/zap"
)

// swagger:model LoginRequestBody
type loginRequest struct {
	// Users username at Telegram
	// in: string
	Username string `json:"username"`
	// Password generated by service
	// in: string
	Password string `json:"password"`
}

type IDGetter interface {
	GetUserID(username string) (int64, error)
}

type AuthRepository interface {
	GetPassword(ctx context.Context, id int64) (string, error)
}

type AuthHandler struct {
	repository AuthRepository
	idGetter   IDGetter
	logger     *zap.SugaredLogger
}
