package auth

import "context"

type loginRequest struct {
	Username string `json:"username"`
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
}

// TODO: Change secret to more secure

var _TEST_SECRET string = `vwoejfnv;np29uwovnp2uiefvbjipb2jcnq`
