package bot

import (
	"context"

	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	client *tele.Bot

	userService      user.UserService
	variablesService user.VariableService
	//pass
}

type HandleFunc func(ctx context.Context, c tele.Context, ch chan error)

type Command struct {
	Meta    tele.Command
	Handler HandleFunc
}
