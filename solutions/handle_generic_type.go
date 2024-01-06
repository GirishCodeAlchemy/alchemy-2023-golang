package solutions

import "fmt"

// Swap is a generic function to swap elements in a slice
func Swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func init() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}

	fmt.Println("Before swap:")
	fmt.Println(intSlice)
	fmt.Println(stringSlice)

	Swap(intSlice, 1, 3)
	Swap(stringSlice, 0, 2)

	fmt.Println("After swap:")
	fmt.Println(intSlice)
	fmt.Println(stringSlice)
}
