package main

import (
	"fmt"
)

func valueFibbonachi(n1 int32) int32 {
	fmt.Scanln(&n1)
	if n1 == 0 {
		return 0
	}

	if n1 == 1 {
		return 1
	}
	return valueFibbonachi(n1-1) + valueFibbonachi(n1-2)
}

func main() {
	fmt.Println("Result: ", valueFibbonachi(5))
}

//не могу понять как можно самому вводить число
