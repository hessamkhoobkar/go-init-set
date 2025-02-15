package math

import (
	"fmt"
)

const a, b = 10, 5

func Calculate() {
	additionSum := Addition(a, b)	
	fmt.Println("Addition:", a, "+", b, "=", additionSum)
	subtractionSum := Subtraction(a, b)
	fmt.Println("Subtraction:", a, "-", b, "=", subtractionSum)
	multiplicationSum := Multiplication(a, b)
	fmt.Println("Multiplication:", a, "*", b, "=", multiplicationSum)
	divisionSum := Division(a, b)
	fmt.Println("Division:", a, "/", b, "=", divisionSum)
	moduloSum := Modulo(a, b)
	fmt.Println("Modulo:", a, "%", b, "=", moduloSum)	
}