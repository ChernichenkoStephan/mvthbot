package bot

import (
	"strings"

	tele "gopkg.in/telebot.v3"
)

func ArgDividerConvertor(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		args := strings.Split(c.Message().Text, "\n")
		c.Set("args", args)
		return next(c)
	}
}
