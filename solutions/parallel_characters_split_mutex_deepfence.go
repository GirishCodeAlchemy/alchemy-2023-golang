package solutions

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var totalcount int = 0
var failurecount int = 0

func worker(finaldata int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	status := checkWithRemoteService(finaldata)
	fmt.Println(status, finaldata)
	m.Lock()
	if status {
		totalcount++
	} else {
		failurecount++
	}
	m.Unlock()
}

func ParallelCharactersSplitMutexDeepfence() {
	fmt.Println("Parallel Characters Split Mutex Deepfence")
	// TODO
	var wg sync.WaitGroup
	var m sync.Mutex
	tmpString := ""

	for {

		b := getNextChar()
		if b == 0 {
			break
		} else {
			tmp := string(b)
			if tmp != ";" {
				_, ok := strconv.Atoi(tmp)
				if ok == nil {
					tmpString += tmp
				}
			} else {
				finaldata, err := strconv.Atoi(tmpString)
				if err != nil {
					print(err)
				}
				fmt.Println(finaldata)
				wg.Add(1)
				go worker(finaldata, &wg, &m)
				tmpString = ""
			}

		}
	}
	wg.Wait()
	fmt.Println(totalcount, failurecount)
}

func init() {

	ParallelCharactersSplitMutexDeepfence()
}

const (
	stream = "12afsd3;345;678;901;234;578"
)

// SKIP BELOW

var (
	index = 0
)

func getNextChar() byte {

	if index < len(stream) {

		tmp := stream[index]

		index += 1

		return tmp

	}

	return 0

}

func checkWithRemoteService(cve_signature int) bool {

	time.Sleep(5 * time.Second)

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1

}
