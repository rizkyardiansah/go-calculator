package go_calculator

import "math"

func Add(a int, b int) int {
	return a + b
}

func Substract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a float64, b float64) float64 {
	return a/b
}

func Modulo(a int, b int) int {
	return a % b
}

func Power(a float64, b float64) float64 {
	return math.Pow(a, b)
}