package main

func Calculations(a int, b int) []float64 {
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	results := []float64{float64(sum), float64(difference), float64(product), float64(quotient)}
	return results
}
