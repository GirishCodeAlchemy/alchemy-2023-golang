package solutions

import (
	"fmt"
	"sync"
	"time"
)

func waitgroupworker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker ", i, "started")
	time.Sleep(time.Second * 2)
	fmt.Println("worker ", i, "finished")
}

func Waitgroupcase() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go waitgroupworker(i, &wg)

	}
	wg.Wait()

}

// func DeferKey() {
// 	fmt.Println("Started defer")
// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}
// 	fmt.Println("Done")
// }

func init() {
	Waitgroupcase()
	// DeferKey()

}
