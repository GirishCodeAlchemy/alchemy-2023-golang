package solutions

import "fmt"

// This function intSeq returns another function, which we define anonymously
// in the body of intSeq. The returned function closes over the variable i to form a closure.

// We call intSeq, assigning the result (a function) to nextInt.
//  This function value captures its own i value, which will be updated each time we call nextInt.

func intseq() func() int {

	i := 0
	return func() int {
		i++
		return i

	}

}

func init() {

	newseq := intseq()
	fmt.Println(newseq())
	fmt.Println(newseq())
	fmt.Println(newseq())
	fmt.Println(newseq())
	newseq = intseq()
	fmt.Println(newseq())
	fmt.Println(newseq())

}
