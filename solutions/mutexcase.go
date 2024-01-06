package solutions

import (
	"fmt"
	"sync"
)

var CommonParam = 0

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	CommonParam += 1
	mutex.Unlock()
	wg.Done()
}

func init() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		w.Add(1)
		go Increment(&w, &m)
	}

	w.Wait()
	fmt.Println("CommonParam: ", CommonParam)
}

func write(CommonMap map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 20; i++ {
			mux.Lock()
			fmt.Println("Writing : ", i)
			CommonMap[i] = i
			mux.Unlock()
		}
	}
}

func read(CommonMap map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for _, value := range CommonMap {
			fmt.Println("Reading value : ", value)
		}
		mux.RUnlock()
	}
}

func init() {
	// CommonMap := map[int]int{}
	// // RWMutex initialization
	// mux := &sync.RWMutex{}

	// go write(CommonMap, mux)
	// go read(CommonMap, mux)
	// go read(CommonMap, mux)

	// //Stop the program from exiting
	// for {
	// }
}
