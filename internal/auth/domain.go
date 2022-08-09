package auth

import "context"

type IDGetter interface {
	GetUserID(username string) (int64, error)
}

type UserAuthRepository interface {
	GetPassword(ctx context.Context, id int64) (string, error)
}

type AuthHandler struct {
	repository UserAuthRepository
	idGetter   IDGetter
}

// TODO: Change secret to more secure

var _TEST_SECRET string = `vwoejfnv;np29uwovnp2uiefvbjipb2jcnq`
