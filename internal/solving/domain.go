package solving

type Statement struct {
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