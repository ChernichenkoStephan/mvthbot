package main

import (
	"fmt"

	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
)

func _test(task string, ref []string) {
	t, err := converting.ToRPN(task)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	fmt.Println(ref)
}

func main() {
	/*
		v, err := converting.ToRPN("123.321+2")
		if err == nil {
			fmt.Println(v)
		}
		res, err := solving.Solve(v)
		if err == nil {
			fmt.Println(res)
		}
	*/

}
