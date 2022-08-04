package converting

// RPNConverter
type RPNConverter interface {
	Convert(line string) ([]string, error)
}

// RPNConverterFunc
type RPNConverterFunc func(line string) ([]string, error)

func ToRPN(line string) []string {
	return []string{}
}
