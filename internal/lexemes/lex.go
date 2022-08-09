package lexemes

const (
	OPERATION LexemeType = iota
	SINGLE_PLACE_FUNC
	DOUBLE_PLACE_FUNC
	SYMBOL
)

var OP_SYM = []rune{'+', '-', '*', '/', '^'}

var FUNC_NAMES = []string{"pow", "log", "mod", "exp", "sqrt"}

var _SUPPORTED_LEXEMES map[string]*Lexeme = map[string]*Lexeme{
	"+": {OPERATION, 0, Sum},
	"-": {OPERATION, 0, Sub},
	"*": {OPERATION, 1, Mult},
	"/": {OPERATION, 1, Div},
	"^": {OPERATION, 2, Pow},

	"(": {SYMBOL, -1, Nothing},
	")": {SYMBOL, -1, Nothing},
	";": {SYMBOL, -2, Nothing},

	"pow": {DOUBLE_PLACE_FUNC, -1, Pow},
	"log": {DOUBLE_PLACE_FUNC, -1, Log},
	"mod": {DOUBLE_PLACE_FUNC, -1, Mod},

	"exp":  {SINGLE_PLACE_FUNC, -1, Exp},
	"sqrt": {SINGLE_PLACE_FUNC, -1, Sqrt},
}

func GetMathOperation(name string) (MathOperation, bool) {
	lex, ok := _SUPPORTED_LEXEMES[name]
	return lex.Op, ok
}

func GetLexType(name string) (LexemeType, bool) {
	lex, ok := _SUPPORTED_LEXEMES[name]
	return lex.LexType, ok
}

func IsLexRune(c rune) bool {
	_, ok := _SUPPORTED_LEXEMES[string(c)]
	return ok
}

func IsLex(name string) bool {
	_, ok := _SUPPORTED_LEXEMES[name]
	return ok
}

func GetLexPriority(name string) (int8, bool) {
	lex, ok := _SUPPORTED_LEXEMES[name]
	return lex.Priority, ok
}
