package converting

import (
	"fmt"
	"strconv"
	"unicode"

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
	if unicode.IsDigit([]rune(name)[0]) {
		return fmt.Errorf("Name can't start with number")
	}
	for _, c := range name {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') && c != '_' {
			return fmt.Errorf("Unvalid variable character: '%c'", c)
		}
	}
	return nil
}

func putVal(buffer []rune, stack []string) ([]rune, []string, error) {
	v := string(buffer)
	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		if err = validateName(v); err != nil {
			return []rune{}, []string{}, fmt.Errorf("Not number or name %v", err)
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
		err     error
	)

	for i, c := range equation {
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
					return []string{}, fmt.Errorf("UnknownFuncError")
				}
			}
		case c == ')':
			if len(buffer) > 0 {
				buffer, output, err = putVal(buffer, output)
				if err != nil {
					return []string{}, fmt.Errorf("Not number %v", err)
				}
			}

			for {
				if len(opStack) == 0 {
					return []string{}, fmt.Errorf("Bracket error %d|%c", i, c)
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
				buffer, output, err = putVal(buffer, output)
				if err != nil {
					return []string{}, fmt.Errorf("Not number %v", err)
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
				buffer, output, err = putVal(buffer, output)
				if err != nil {
					return []string{}, fmt.Errorf("Not number %v", err)
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
		buffer, output, err = putVal(buffer, output)
		if err != nil {
			return []string{}, fmt.Errorf("Not number %v", err)
		}
	}

	for {
		if len(opStack) == 0 {
			break
		}

		op, opStack = utils.Pop(opStack)
		if op == "(" || op == ")" {
			return []string{}, fmt.Errorf("Bracket error")
		}
		output = append(output, op)
	}

	return output, nil
}
