package bot

import (
	"strings"

	slv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
	tele "gopkg.in/telebot.v3"
)

// TODO make stack based more efficient custom function
func parseStatements(input string) []slv.Statement {
	statements := make([]slv.Statement, 0)

	// For multiline input
	lines := strings.Split(input, "\n")

	// End of command
	comInd := strings.IndexAny(lines[0], " ")

	// For no arg command
	if comInd == -1 {
		return statements
	}

	// Removing command from first line
	lines[0] = lines[0][comInd:]

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
