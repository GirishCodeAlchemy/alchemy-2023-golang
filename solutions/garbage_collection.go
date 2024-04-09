package solutions

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC, "\n")
}

// The runtime.ReadMemStats() call gets the latest garbage collection statistics

func init() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		// Allocating 50,000,000 bytes
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)
	// 	In this part, we have a for loop that creates 10-byte slices with 50,000,000 bytes each.
	// The reason for this is that by allocating large amounts of memory, we can trigger the GC.

	for i := 0; i < 10; i++ {
		// Allocating 100,000,000 bytes
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}

// The value of mem.Alloc is the bytes of allocated heap objects — allocated are all the objects that the GC has not yet freed. mem.TotalAlloc shows the cumulative bytes allocated for heap objects—this number does not decrease when objects are freed, which means that it keeps increasing. Therefore, it shows the total number of bytes allocated for heap objects during program execution. mem.HeapAlloc is the same as mem.Alloc. Last, mem.NumGC shows the total number of completed garbage collection cycles. The bigger that value is, the more you have to consider how you allocate memory in your code and if there is a way to optimize that.

// If you want even more verbose output regarding the operation of the GC, you can
// combine go run gColl.go with GODEBUG=gctrace=1. Apart from the regular program
// output, you get some extra metrics—this is illustrated in the following output:

// https://hub.packtpub.com/implementing-memory-management-with-golang-garbage-collector/

// https://github.com/golang/go/blob/master/src/runtime/mgc.go,
