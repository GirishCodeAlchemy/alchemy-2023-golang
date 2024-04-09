package main

import (
	lib "GoLearning/lib"
	"fmt"
)

const PI = 3.14
const (
	A = iota
	B = iota
)

func main() {
	println("Hi this is first code")
	fmt.Printf("This is printf %s", "string")
	var first int
	first = 10
	fmt.Println(first)
	var second string = "This is string"
	fmt.Println(second)
	// Raw string literals
	third := `This is
	third"`
	fmt.Println(third)
	var fourth, fifth int = 10, 20
	fmt.Println(fourth, fifth)
	var sixth, seventh = "This is string", 20
	fmt.Println(sixth, seventh)
	var (
		name = "go"
		age  = 10
	)
	fmt.Println(name, age)

	// constant
	fmt.Println(A, B)

	// Type assertion
	var ij interface{} = "gola"
	t, isSuccess := ij.(string)
	fmt.Println(t, isSuccess)

	// array
	var arr [3]string
	arr[0] = "go"
	arr[1] = "java"
	fmt.Println(arr, arr[2])

	cities := [3]string{"go", "java", "python"}
	fmt.Println(cities)
	fmt.Println(cities[1:])

	cities1 := [...]string{"go", "java", "python"}
	fmt.Println(cities1, len(cities1))
	// fmt.Println(cities1[-1])    // error

	// slices
	var cities3 []string
	cities3 = append(cities3, "go")
	fmt.Println(cities3)
	fmt.Printf("cities3 : %v\n", cities3)

	var cities4 = make([]string, 3)
	fmt.Println(cities4, len(cities4))

	fmt.Println(cities3[1:])
	// infinte slice
	list := make([]string, 0)
	list = append(list, "go")
	fmt.Println(list)

	// Iteration
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	for i, v := range cities {
		fmt.Println(i, v)
	}

	var ctr int = 0
	for ctr < 5 {
		fmt.Println(ctr)
		ctr++
	}

	// map
	employees := make(map[string]int)
	employees["Mark"] = 10
	employees["Sandy"] = 20
	delete(employees, "Sandy")
	fmt.Println(employees)

	employyes2 := map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employyes2)

	// swithc
	lib.SwitchCaseFunction()
	lib.SwitchCaseLabelFunction()

	// // structure
	// lib.StructureFunction()
	// emp5 := lib.Employee{Name: "giri", Age: 25}
	// emp5.Display()
	// emp6 := lib.EmployeePointers{Name: "giri", Age: 25}
	// emp6.Display()
	// emp6.Display()

	// // Interface
	// lib.InterfaceFunction()

	// // // Goroutines
	// // lib.GoroutinesFunction()

	// // // Channels
	// // lib.ChannelFunctions()

	// // // webserver
	// // lib.SimpleWebserver()

	// //StingsFunction
	// lib.StingFunction()

	// // Recursive
	// fmt.Println("factorial of 5", lib.Recursive(5))
	// fmt.Println("Fibanocii of 10", lib.Fibonacci(9))

	// // Error
	// lib.ErrorCase()

	// // SUbstring
	// lib.SubstringFunction("girish")

	// // Solutions
	// // 1. custom sort
	// sol.CustomSort()

}
