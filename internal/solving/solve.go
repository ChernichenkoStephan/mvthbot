package solving

import (
	"fmt"
	"strconv"

	lex "github.com/ChernichenkoStephan/mvthbot/internal/lexemes"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"
)

func Solve(equation []string, variables VMap) (float64, error) {
	stack := []float64{}
	var l, r float64
	var ok bool

	for _, e := range equation {
		if lex.IsLex(e) {
			r, stack = utils.Pop(stack)
			f, _ := lex.GetMathOperation(e)

			switch t, _ := lex.GetLexType(e); t {
			case lex.SINGLE_PLACE_FUNC:
				v, err := f(r)
				if err != nil {
					return 0.0, fmt.Errorf("WrongFunctionUsage")
				}
				stack = append(stack, v)

			case lex.DOUBLE_PLACE_FUNC, lex.OPERATION:
				l, stack = utils.Pop(stack)
				v, err := f(l, r)
				if err != nil {
					return 0.0, fmt.Errorf("WrongFunctionUsage")
				}
				stack = append(stack, v)

			}
		} else {
			v, err := strconv.ParseFloat(e, 64)
			if err != nil {
				v, ok = variables[e]
				if !ok {
					return 0.0, fmt.Errorf("Unknown variable %s", e)
				}
			}
			stack = append(stack, v)
		}
	}

	if len(stack) == 1 {
		return stack[0], nil
	}

	return 0.0, fmt.Errorf("Uncorrect equation")
}
