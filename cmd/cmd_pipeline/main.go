package main

import (
	"fmt"

	calculator "github.com/ErickRoa95/hello-actions/internal"
)

func main() {
	calc := calculator.NewCalculator()
	a, b := 10, 5

	fmt.Printf("%d + %d = %d\n", a, b, calc.Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, calc.Subtract(a, b))
	fmt.Printf("%d * %d = %.2f\n", a, b, calc.Multiply(a, b))
	fmt.Printf("%d / %d = %.2f\n", a, b, calc.Divide(a, b))
}