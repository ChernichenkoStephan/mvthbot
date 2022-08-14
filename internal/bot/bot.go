package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
	"github.com/ChernichenkoStephan/mvthbot/internal/fixing"
	slv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	tele "gopkg.in/telebot.v3"
)

func NewBot(
	client *tele.Bot,
	userService user.UserService,
	variablesService user.VariableService,
	fixer fixing.Fixer,
) *Bot {
	return &Bot{
		client:           client,
		userService:      userService,
		variablesService: variablesService,
		stringFixer:      fixer,
	}
}

func (b Bot) Client() *tele.Bot {
	return b.client
}

func (b *Bot) GetUserID(username string) (int64, error) {
	return 11111, nil
}

func (b *Bot) Broadcast(message string) error {
	ctx := context.TODO()
	users, err := b.userService.GetAll(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < len(*users); i++ {
		go func(u *user.User) {
			b.client.Send(u, message)
		}(&(*users)[i])
	}

	return nil
}

func (b *Bot) process(ctx context.Context, uID int64, statements interface{}) (string, error) {
	sts, ok := statements.([]slv.Statement)
	if !ok {
		return "", fmt.Errorf("InputError %v", sts)
	}

	builder := NewOutputBuilder()

	vs, err := b.variablesService.GetAll(context.TODO(), uID)
	if err != nil {
		return "", err
	}

	for _, s := range sts {

		fixed := b.stringFixer.Fix(s.Equation)

		eq, err := converting.ToRPN(fixed)
		if err != nil {
			return "", err
		}
		res, err := slv.Solve(eq, vs)
		if err != nil {
			return "", err
		}
		s.Value = res

		c := context.TODO()

		err = b.variablesService.AddWithNames(c, uID, s.Variables, res)
		if err != nil {
			return "", err
		}
		builder.Write(&s)

		c = context.TODO()
		err = b.userService.AddStatement(c, uID, &s)
		if err != nil {
			log.Println("History save error")
		}

		for _, n := range s.Variables {
			vs[n] = res
		}

	}
	return builder.String(), nil
}
