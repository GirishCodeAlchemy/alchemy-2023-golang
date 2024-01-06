package lib

import (
	"fmt"
	"strings"
)

func StingFunction() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.HasPrefix("chicken", "ch"))
	fmt.Println(strings.HasSuffix("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToUpper("gopher"))
	fmt.Println(strings.Trim(" !!! Agfsdchtung !!! ", "! "))
	fmt.Println(strings.TrimLeft(" !!! Achtung !!! ", "! "))
	fmt.Println(strings.TrimRight(" !!! Achtung !!! ", "! "))
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n "))
	fmt.Println(strings.Join([]string{"a", "b"}, ","))
	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))
	// fmt.Println(strings.ContainsAny(

}
