package bot

import (
	"context"
	"strings"

	"emperror.dev/errors"
	slv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/user"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

// TODO make stack based more efficient custom function
func parseStatements(input string) []slv.Statement {
	statements := make([]slv.Statement, 0)

	// For multiline input
	lines := strings.Split(input, "\n")

	// End of command
	comInd := strings.IndexAny(lines[0], " ")

	// If it is command
	if lines[0][:1] == "/" {
		// For no arg command
		if comInd == -1 {
			return statements
		}

		// Removing command from first line
		lines[0] = lines[0][comInd:]
	}

	// Parsing all lines
	var vars []string
	for _, l := range lines {
		cleared := strings.ReplaceAll(l, " ", "")
		parts := strings.Split(cleared, "=")

		// Getting equation from last part
		eq := parts[len(parts)-1]

		// If there is no vars, just equation
		if len(parts) == 1 {
			vars = []string{}
		} else {
			vars = parts[:len(parts)-1]
		}
		statements = append(statements, slv.Statement{
			Variables: vars,
			Equation:  eq,
		})

	}

	return statements
}

func ArgParse(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		sts := parseStatements(c.Message().Text)
		c.Set("statements", sts)
		return next(c)
	}
}

func Logging(logger *zap.SugaredLogger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			logger.Infof("Message from '%v' with text '%v' in %s[%s|%d]", c.Sender().ID, c.Text(), c.Chat().FirstName, c.Chat().Type, c.Chat().ID)
			return next(c)
		}
	}
}

func UserCheck(db *user.Database) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat().Type != "channel" {
				ctx := context.TODO()
				ok, err := db.Exist(ctx, c.Message().Sender.ID)
				if err != nil {
					return errors.Wrap(err, "User existence check failed")
				}
				if !ok {
					u := user.NewUser(c.Message().Sender.ID)
					db.Add(ctx, u.TelegramID, u.Password)
				}
			}
			return next(c)
		}
	}

}
