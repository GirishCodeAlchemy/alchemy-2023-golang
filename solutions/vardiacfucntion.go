package solutions

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total

}

func VardiacFunction() {
	fmt.Println("Vardiac Function")
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}
