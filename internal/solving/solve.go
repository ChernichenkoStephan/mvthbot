package solving

import (
	"fmt"
	"strconv"

	lex "github.com/ChernichenkoStephan/mvthbot/internal/lexemes"
)

func pop(stack []float64) (float64, []float64) {
	op := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return op, stack
}

func Solve(equation []string) (float64, error) {
	stack := []float64{}
	var l, r float64

	for _, e := range equation {
		if lex.IsLex(e) {
			r, stack = pop(stack)
			f, _ := lex.GetMathOperation(e)

			switch t, _ := lex.GetLexType(e); t {
			case lex.SINGLE_PLACE_FUNC:
				v, err := f(r)
				if err != nil {
					return 0.0, fmt.Errorf("WrongFunctionUsage")
				}
				stack = append(stack, v)

			case lex.DOUBLE_PLACE_FUNC, lex.OPERATION:
				l, stack = pop(stack)
				v, err := f(l, r)
				if err != nil {
					return 0.0, fmt.Errorf("WrongFunctionUsage")
				}
				stack = append(stack, v)

			}
		} else {
			v, err := strconv.ParseFloat(e, 64)
			if err != nil {
				return 0.0, fmt.Errorf("Not number %v", err)
			}
			stack = append(stack, v)
		}
	}

	if len(stack) == 1 {
		return stack[0], nil
	}

	return 0.0, fmt.Errorf("Uncorrect equation")
}
