package fixing

import "strings"

var _ERROR_CASES map[string]string = map[string]string{
	"+-": "-",
	"-+": "-",
	"--": "+",
	"++": "+",
	"(+": "(",
	"(-": "(0-",
}

type FixFunc func(line string) string

type Fixer interface {
	Fix(line string) string
}

type BaseFixer struct {
	fixes []FixFunc
}

func (f *BaseFixer) Handle(fix FixFunc) {
	f.fixes = append(f.fixes, fix)
}

func (f BaseFixer) Fix(equation string) string {
	res := equation
	for _, fx := range f.fixes {
		res = fx(res)
	}
	return res
}

func New() *BaseFixer {
	return &BaseFixer{[]FixFunc{fixSpaces, fixErrors}}
}

func fixSpaces(equation string) string {
	return strings.ReplaceAll(equation, " ", "")
}

func fixErrors(equation string) string {
	res := equation
	for e, c := range _ERROR_CASES {
		res = strings.ReplaceAll(res, e, c)
	}
	return res
}
