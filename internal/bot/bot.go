package bot

import (
	"context"
	"runtime"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
	"github.com/ChernichenkoStephan/mvthbot/internal/fixing"
	slv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	tele "gopkg.in/telebot.v3"
)

func NewBot(
	client *tele.Bot,
	db *user.Database,
	fixer fixing.Fixer,
	lg *zap.SugaredLogger,
	config *BotConfig,
) *Bot {
	return &Bot{
		client:      client,
		db:          db,
		stringFixer: fixer,
		logger:      lg,
		conf:        config,
	}
}

func (b Bot) Client() *tele.Bot {
	return b.client
}

func (b Bot) GetCommandText(command string) string {
	if txt, ok := b.ComandsTexts[command]; ok {
		return txt
	}
	return `Text not set (ಥ﹏ಥ)`
}

func (b Bot) SetCommandText(command, text string) {
	b.ComandsTexts[command] = text
}

func (b *Bot) GetUserID(username string) (int64, error) {
	c, err := b.client.ChatByUsername(username)
	if err != nil {
		return 0, errors.Wrap(err, `chat fetch failed`)
	}
	return c.ID, nil
}

func (b *Bot) Broadcast(ctx context.Context, message string) error {
	users, err := b.db.GetAll(ctx)
	if err != nil {
		return err
	}

	_, err = b.client.Send((*users)[0], message)
	if err != nil {
		return errors.Wrap(err, `Brodcasting failed`)
	}

	group, _ := errgroup.WithContext(ctx)

	if len(*users) > 1 {
		targets := (*users)[1:]
		batchamm := runtime.NumCPU()
		batchlen := len(targets) / batchamm

		for i := 0; i < len(*users); i += batchlen {
			start := i
			end := batchlen
			batch := (*users)[start:end]
			group.Go(func() error {
				for j := start; j < end; j++ {
					_, err := b.client.Send(batch[j], message)
					if err != nil {
						b.logger.Errorf("Brodcast send error. To user %v, with error: %s", batch[j].TelegramID, err.Error())
					}
					if i == 0 {
						return errors.Wrap(err, `Brodcasting failed`)
					}
				}
				return nil
			})
		}

	}

	return group.Wait()
}

func (b *Bot) process(ctx context.Context, uID int64, sts []slv.Statement) (string, error) {
	builder := NewOutputBuilder()

	err := b.db.WithinTransaction(ctx, func(ctx context.Context) error {
		vs, err := b.db.GetAllVariables(ctx, uID)
		if err != nil {
			msg := "DB get all variables failed"
			return errors.Wrap(err, msg)
		}

		for _, s := range sts {

			fixed := b.stringFixer.Fix(s.Equation)

			eq, err := converting.ToRPN(fixed)
			if err != nil {
				msg := "Converting to RPN failed"
				return errors.Wrap(err, msg)
			}
			res, err := slv.Solve(eq, vs)
			if err != nil {
				msg := "Solving failed"
				return errors.Wrap(err, msg)
			}
			s.Value = res

			builder.Write(&s)

			err = b.db.AddStatement(ctx, uID, &s)
			if err != nil {
				msg := "Statements add failed"
				return errors.Wrap(err, msg)
			}

			for _, n := range s.Variables {
				vs[n] = res
			}

		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
