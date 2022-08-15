package solving

import (
	"fmt"
	"strconv"

	"emperror.dev/errors"
	lex "github.com/ChernichenkoStephan/mvthbot/internal/lexemes"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"
)

func Solve(equation []string, variables map[string]float64) (float64, error) {
	stack := []float64{}
	var l, r float64
	var ok bool

	for _, element := range equation {
		if lex.IsLex(element) {
			r, stack = utils.Pop(stack)
			f, _ := lex.GetMathOperation(element)

			switch t, _ := lex.GetLexType(element); t {
			case lex.SINGLE_PLACE_FUNC:
				v, err := f(r)
				if err != nil {
					msg := fmt.Sprintf("Wrong %s usage with param %f", element, r)
					return 0.0, errors.Wrap(err, msg)
				}
				stack = append(stack, v)

			case lex.DOUBLE_PLACE_FUNC, lex.OPERATION:
				l, stack = utils.Pop(stack)
				v, err := f(l, r)
				if err != nil {
					msg := fmt.Sprintf("Wrong %s usage with params %f, %f", element, l, r)
					return 0.0, errors.Wrap(err, msg)
				}
				stack = append(stack, v)

			}
		} else {
			v, err := strconv.ParseFloat(element, 64)
			if err != nil {
				v, ok = variables[element]
				if !ok {
					return 0.0, fmt.Errorf("Unknown variable %s", element)
				}
			}
			stack = append(stack, v)
		}
	}

	if len(stack) != 1 {
		return 0.0, fmt.Errorf("Uncorrect equation | stack %v", stack)
	}
	return stack[0], nil

}
