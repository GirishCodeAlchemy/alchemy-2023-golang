// Package library
package library

func Recursive(i int) int {
	if i <= 1 {
		return 1
	}
	// fmt.Println(i, i-1)
	return i * Recursive(i-1)
}

func Fibonacci(i int) int {
	if i <= 1 {
		return i
	}
	// fmt.Println(i-1, i-2)
	return Fibonacci(i-1) + Fibonacci(i-2)
}
