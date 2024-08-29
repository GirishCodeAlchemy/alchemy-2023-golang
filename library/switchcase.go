package library

import (
	"fmt"
)

func SwitchCaseFunction() {
	var a int = 10
	switch a {
	case 5:
		fmt.Println("a is 5")
		fallthrough
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("a is not 5 or 10")
	}
}

func SwitchCaseLabelFunction() {
myloop:
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ,%T ", i, i)
		switch {
		case i == 5:
			fmt.Println("i is 5")
			break myloop
		case i == 7:
			fmt.Println("i is 7")
		default:
			fmt.Println("i is not 5 or 7")
		}
	}
	var a int = 10
	switch a {
	case 5:
		fmt.Println("a is 5")
		fallthrough
	case 10:
		fmt.Println("a is 10")
	default:
		fmt.Println("a is not 5 or 10")
	}
}

// func init() {
// 	SwitchCaseFunction()
// }

func Swithccasetype() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}
}
