package solving

type Statement struct {

	// Internal DB id
	Id int

	// variables to set
	// example: a = b = 2 + 2 >> Variables = ["a", "b"]
	Variables []string

	// Equation to solve
	// example: a = b = 2 + 2 >> Equation = "2 + 2"
	Equation string

	// Equation result value
	// example: a = b = 2 + 2 >> Value = 4.0
	Value float64
}

//
//
// DTO for views
//
//

// swagger:model Statement
type PackDTO struct {
	// Equations for solve
	// Example: ["2+2", "3+3", "4+4"]
	// in: string[]
	Equations []string `validate:"required"`
}
