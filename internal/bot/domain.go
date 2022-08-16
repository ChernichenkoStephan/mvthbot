package bot

import (
	"context"

	"github.com/ChernichenkoStephan/mvthbot/internal/fixing"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {

	// Telegram Bot API client
	client *tele.Bot

	// Working with users personal data
	userService user.UserService

	// Workiing with users variables
	variablesService user.VariableService

	// Fixing mistakes in input
	stringFixer fixing.Fixer

	// log Logger
	logger *zap.SugaredLogger
}

type HandleFunc func(ctx context.Context, c tele.Context) error

type Command struct {
	Meta            tele.Command
	Handler         HandleFunc
	IsParameterized bool
}
