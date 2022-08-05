package converting

import (
	"fmt"
	"strconv"
	"unicode"

	lex "github.com/ChernichenkoStephan/mvthbot/internal/lexemes"
)

// RPNConverter
type RPNConverter interface {
	Convert(line string) ([]string, error)
}

// RPNConverterFunc
type RPNConverterFunc func(line string) ([]string, error)

func putNum(buffer []rune, stack []string) ([]rune, []string, error) {
	num := string(buffer)
	_, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return []rune{}, []string{}, fmt.Errorf("Not number %v", err)
	}
	stack = append(stack, num)
	buffer = []rune{}
	return buffer, stack, nil
}

func pop(stack []string) (string, []string) {
	op := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return op, stack
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
				buffer, output, err = putNum(buffer, output)
				if err != nil {
					return []string{}, fmt.Errorf("Not number %v", err)
				}
			}

			for {
				if len(opStack) == 0 {
					return []string{}, fmt.Errorf("Bracket error %d|%c", i, c)
				}

				op, opStack = pop(opStack)
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
				buffer, output, err = putNum(buffer, output)
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
				op, opStack = pop(opStack)
				output = append(output, op)
			}
		case lex.IsLexRune(c):
			if len(buffer) != 0 {
				buffer, output, err = putNum(buffer, output)
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
					op, opStack = pop(opStack)
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
		buffer, output, err = putNum(buffer, output)
		if err != nil {
			return []string{}, fmt.Errorf("Not number %v", err)
		}
	}

	for {
		if len(opStack) == 0 {
			break
		}

		op, opStack = pop(opStack)
		if op == "(" || op == ")" {
			return []string{}, fmt.Errorf("Bracket error")
		}
		output = append(output, op)
	}

	return output, nil
}
