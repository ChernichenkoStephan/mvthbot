package lexemes

import (
	"errors"
	"fmt"
	"math"
)

func Nothing(args ...float64) (float64, error) {
	fmt.Println("Nothing")
	return 0.0, nil
}

func Sum(args ...float64) (float64, error) {
	return args[0] + args[1], nil
}

func Sub(args ...float64) (float64, error) {
	return args[0] - args[1], nil
}

func Mult(args ...float64) (float64, error) {
	return args[0] * args[1], nil
}

func Div(args ...float64) (float64, error) {
	if args[1] < 0.000001 {
		return 0.0, errors.New("Divide by 0")
	}
	return args[0] / args[1], nil
}

func Pow(args ...float64) (float64, error) {
	return math.Pow(args[0], args[1]), nil
}

// Log вычисляет логарифм с
// основанием > 1 и x более 0
func Log(args ...float64) (float64, error) {
	return math.Log(args[0]) / math.Log(args[1]), nil
}

func Mod(args ...float64) (float64, error) {
	// TODO
	return 0.0, nil
}

func Exp(args ...float64) (float64, error) {
	return math.Exp(args[0]), nil
}

func Sqrt(args ...float64) (float64, error) {
	return math.Sqrt(args[0]), nil
}
