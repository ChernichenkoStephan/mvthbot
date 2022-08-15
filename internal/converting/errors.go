package converting

import "fmt"

// , UnknownFuncError

type VariableNameError struct {
	Name      string
	Msg       string
	Character rune
	Place     int
}

func (e VariableNameError) Error() string {
	return fmt.Sprintf("%s | name: %s | %d-nth char '%c'", e.Msg, e.Name, e.Place, e.Character)
}

type UnknownValueError struct {
	Value string
	Place int
}

func (e UnknownValueError) Error() string {
	return fmt.Sprintf("Unknown value '%s' at %d", e.Value, e.Place)
}

type UnknownFunctionError struct {
	Name  string
	Place int
}

func (e UnknownFunctionError) Error() string {
	return fmt.Sprintf("Unknown function '%s' at %d", e.Name, e.Place)
}

type BracketError struct {
	Place int
}

func (e BracketError) Error() string {
	return fmt.Sprintf("Wrong bracket usage at %d", e.Place)
}
