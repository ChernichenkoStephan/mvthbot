package lexemes

type LexemeType int

type Lexeme struct {
	LexType  LexemeType
	Priority int8
	Op       MathOperation
}

type MathOperation func(args ...float64) (float64, error)
