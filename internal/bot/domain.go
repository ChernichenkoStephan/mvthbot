package bot

import (
	"context"

	"github.com/ChernichenkoStephan/mvthbot/internal/fixing"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type BotConfig struct {
	PasswordLength int
	AdminKey       string
}

type Bot struct {

	// Telegram Bot API client
	client *tele.Bot

	// Working with users personal data
	db *user.Database

	// Fixing mistakes in input
	stringFixer fixing.Fixer

	// Log Logger
	logger *zap.SugaredLogger

	// Text to output
	ComandsTexts map[string]string

	conf *BotConfig
}

type HandleFunc func(ctx context.Context, c tele.Context) error

type Command struct {
	Meta      tele.Command
	Handler   HandleFunc
	IsPrivate bool
}
