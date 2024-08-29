package library

import "fmt"

func add(a int, b int) {
	fmt.Println("---->", a, b, "-->", a+b)
}

func GoroutinesFunction() {
	for i := 0; i < 10; i++ {
		go add(i, i)
	}
	fmt.Println("---->done")
	add(30, 20)
	var abc string

	go func() {
		fmt.Println("Annonynmous goroutine")
	}()

	fmt.Scanln(&abc)
	fmt.Println("---->", abc)
	fmt.Println("--->")
}
