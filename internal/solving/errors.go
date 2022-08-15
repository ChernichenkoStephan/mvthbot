package solving

import (
	"fmt"
	"strings"
)

type MultySolveError struct {
	Errors []string
}

func (e MultySolveError) Error() string {
	var buffer strings.Builder
	for _, err := range e.Errors {
		buffer.WriteString(err)
		buffer.WriteString("\n")
	}
	return fmt.Sprintf("Got errors during proccessing. Errors:\n%s", buffer.String())
}
