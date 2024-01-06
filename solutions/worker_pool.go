package solutions

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(2 * time.Second)
		results <- j

	}
	fmt.Println("")

}
func WorkerPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 10; a++ {
		fmt.Println("Finished job", <-results)
	}

}

func init() {
	WorkerPool()
}
