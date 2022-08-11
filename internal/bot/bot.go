package bot

import (
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	tele "gopkg.in/telebot.v3"
)

func NewBot(
	client *tele.Bot,
	userService user.UserService,
	variablesService user.VariableService,
) *Bot {
	return &Bot{
		client:           client,
		userService:      userService,
		variablesService: variablesService,
	}
}

func (b Bot) Client() *tele.Bot {
	return b.client
}

func (b *Bot) GetUserID(username string) (int64, error) {
	return 11111, nil
}
