package library

import (
	"fmt"
	"math"
)

// RangeForLoop ...
func RangeForLoop() {
	// pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	pow0 := []int{0, 1, 2, 3, 4, 5, 6}
	for i, v := range pow0 {
		fmt.Printf("2**%d = %d\n", i, int(math.Pow(2, float64(v))))
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
func ForLoopIterator() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func WhileForLoop() {
	var ctr int = 0
	for ctr < 5 {
		fmt.Println(ctr)
		ctr++
	}
}

func init() {
	fmt.Println("\n******** For Loop smaple code *********")
	ForLoopIterator()
	RangeForLoop()
	WhileForLoop()
}
