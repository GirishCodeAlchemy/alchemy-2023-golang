package lib

import (
	"fmt"
	"slices"
)

func SliceFunctions() {
	fmt.Println("Slice Functions")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Contains(slice, 3))
	fmt.Println(slices.Index(slice, 3))
	fmt.Println(slices.Max(slice))
	fmt.Println(slices.Min(slice))
	// a := slices.SortFunc(slice, func(a, b int) bool {
	// 	return a < b
	// })
	// fmt.Println(slices.SortFunc(slice, func(a, b int) bool {
	// 	return a > b
	// }))
	// fmt.Println(slices.SortFunc(slice, func(a, b int) bool {
	// 	return a == b
	// }))
	names := []string{"alice", "Bob", "VERA"}
	fmt.Println(slices.Contains(names, "alice"))
	slices.Reverse(names)
	fmt.Println(names)
	slices.Sort(names)
	fmt.Println(names)
}

func init() {
	SliceFunctions()
}
