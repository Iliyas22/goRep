package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Welcome to calculator:")

	var a, b, total float64
	var operation string

	fmt.Println("Enter a first value:")
	fmt.Scanln(&a)

	fmt.Println("Choose operation:+,-,/,*,^,%")
	fmt.Scanln(&operation)

	fmt.Println("Enter a second value:")
	fmt.Scanln(&b)

	switch operation {
	case "+":
		{
			total = a + b
			fmt.Println("Output: ", +total)
		}
	case "-":
		{
			total = a - b
			fmt.Println("Output: ", +total)
		}
	case "/":
		{
			if b == 0 {
				fmt.Println("Output: Error! You cannot divide a number by 0!")
			} else {
				total = a / b
				fmt.Println("Output: ", +total)
			}
		}
	case "*":
		{
			total = a * b
			fmt.Println("Output: ", +total)
		}
	case "%":
		{
			total = (b * a) / 100
			fmt.Println("Output: ", b, "%"+" of the number", +a, "=", +total)
		}
	case "^":
		{
			total = math.Pow(a, b)
			fmt.Println("Output: ", +total)
		}

	default:
		{
			fmt.Println("Error operation!")
			os.Exit(1)
		}

	}

}
