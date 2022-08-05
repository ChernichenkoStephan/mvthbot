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

type Fixer struct {
	fixes []FixFunc
}

func (f *Fixer) Handle(fix FixFunc) {
	f.fixes = append(f.fixes, fix)
}

func (f Fixer) Fix(equation string) string {
	res := equation
	for _, fx := range f.fixes {
		res = fx(res)
	}
	return res
}

func New() *Fixer {
	return &Fixer{[]FixFunc{fixSpaces, fixErrors}}
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
