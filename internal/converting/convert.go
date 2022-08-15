package converting

import (
	"fmt"
	"strconv"
	"unicode"

	"emperror.dev/errors"
	lex "github.com/ChernichenkoStephan/mvthbot/internal/lexemes"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"
)

// RPNConverter
type RPNConverter interface {
	Convert(line string) ([]string, error)
}

// RPNConverterFunc
type RPNConverterFunc func(line string) ([]string, error)

func validateName(name string) error {
	begin := []rune(name)[0]
	if unicode.IsDigit(begin) {
		return VariableNameError{"Name can't start with number", name, begin, 0}
	}
	for i, c := range name {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') && c != '_' {
			return VariableNameError{"Name can't start with number", name, c, i}
		}
	}
	return nil
}

func putVal(buffer []rune, stack []string) ([]rune, []string, error) {
	v := string(buffer)
	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		if err = validateName(v); err != nil {
			return []rune{}, []string{}, errors.Wrap(err, "Put value to stack failed")
		}
	}
	stack = append(stack, v)
	buffer = []rune{}
	return buffer, stack, nil
}

func ToRPN(equation string) ([]string, error) {

	var (
		output  = []string{}
		opStack = []string{}
		buffer  = []rune{}
		op      string
	)

	for place, c := range equation {
		switch {
		case unicode.IsDigit(c) || c == '.':
			buffer = append(buffer, c)
		case c == '(':
			// if it is open bracket
			if len(buffer) == 0 {
				opStack = append(opStack, string(c))

			} else {
				name := string(buffer)
				if t, ok := lex.GetLexType(name); ok {
					if t == lex.SINGLE_PLACE_FUNC || t == lex.DOUBLE_PLACE_FUNC {
						opStack = append(opStack, name)
						buffer = []rune{}
					}
				} else {
					return []string{}, UnknownFunctionError{name, place}
				}
			}

		case c == ')':
			if len(buffer) > 0 {
				var err error
				buffer, output, err = putVal(buffer, output)
				if err != nil {
					msg := fmt.Sprintf("Error during parsing at %d char: '%c'", place, c)
					return []string{}, errors.Wrap(err, msg)
				}
			}

			for {
				if len(opStack) == 0 {
					return []string{}, BracketError{place}
				}

				op, opStack = utils.Pop(opStack)
				if t, ok := lex.GetLexType(op); ok {
					if t == lex.SINGLE_PLACE_FUNC || t == lex.DOUBLE_PLACE_FUNC {
						output = append(output, op)
						break
					}
					if op == "(" {
						break
					}
					output = append(output, op)
				}
			}

		case c == ';':
			if len(buffer) != 0 {
				var err error
				buffer, output, err = putVal(buffer, output)
				if err != nil {
					msg := fmt.Sprintf("Error during parsing at %d char: '%c'", place, c)
					return []string{}, errors.Wrap(err, msg)
				}
			}

			for {
				if len(opStack) == 0 {
					break
				}

				t, _ := lex.GetLexType(opStack[len(opStack)-1])
				if t == lex.SINGLE_PLACE_FUNC || t == lex.DOUBLE_PLACE_FUNC {
					break
				}
				op, opStack = utils.Pop(opStack)
				output = append(output, op)
			}

		case lex.IsLexRune(c):
			if len(buffer) != 0 {
				var err error

				buffer, output, err = putVal(buffer, output)
				if err != nil {
					msg := fmt.Sprintf("Error during parsing at %d char: '%c'", place, c)
					return []string{}, errors.Wrap(err, msg)
				}

			}

			for {
				if len(opStack) == 0 {
					break
				}

				hp, _ := lex.GetLexPriority(opStack[len(opStack)-1])
				np, _ := lex.GetLexPriority(string(c))

				if hp >= np {
					op, opStack = utils.Pop(opStack)
					output = append(output, op)
				} else {
					break
				}
			}
			opStack = append(opStack, string(c))

		default:
			buffer = append(buffer, c)
		}
	}

	if len(buffer) > 0 {
		var err error
		buffer, output, err = putVal(buffer, output)
		if err != nil {
			buffer, output, err = putVal(buffer, output)
			msg := fmt.Sprintf("Error during parsing")
			return []string{}, errors.Wrap(err, msg)
		}
	}

	for {
		if len(opStack) == 0 {
			break
		}

		op, opStack = utils.Pop(opStack)
		if op == "(" || op == ")" {
			return []string{}, BracketError{len(equation)}
		}
		output = append(output, op)
	}

	return output, nil
}
