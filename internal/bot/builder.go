package bot

import (
	"fmt"

	slv "github.com/ChernichenkoStephan/mvthbot/internal/solving"
)

type OutputBuilder struct {
	data []byte
}

func (b *OutputBuilder) WriteEquation(equation string) {
	b.data = append(b.data, []byte(equation+" = ")...)
}

func (b *OutputBuilder) WriteVariable(variable string) {
	b.data = append(b.data, []byte(variable+" = ")...)
}

func (b *OutputBuilder) WriteVariables(variables *[]string) {
	if len(*variables) == 0 {
		b.WriteVariable("_")
	}
	for _, v := range *variables {
		b.WriteVariable(v)
	}
}

func (b *OutputBuilder) WriteValue(value float64) {
	b.data = append(b.data, []byte(fmt.Sprintf("%v\n", value))...)
}

func (b *OutputBuilder) Write(statement *slv.Statement) {
	b.WriteVariables(&statement.Variables)
	b.WriteValue(statement.Value)
}

func (b *OutputBuilder) WriteFull(statement *slv.Statement) {
	b.WriteVariables(&statement.Variables)
	b.WriteEquation(statement.Equation)
	b.WriteValue(statement.Value)
}

func (b *OutputBuilder) LineBreak() {
	b.data = append(b.data, []byte("\n")...)
}

func (b *OutputBuilder) String() string {
	return string(b.data)
}

func NewOutputBuilder() *OutputBuilder {
	return &OutputBuilder{make([]byte, 0)}
}
