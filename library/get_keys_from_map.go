package library

import (
	"fmt"
	"reflect"
	"strings"
)

func init() {
	fmt.Println("\n******** Get Key maps *********")

	a := map[string]int{
		"A": 1, "B": 2,
	}
	keys := reflect.ValueOf(a).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	fmt.Println(strkeys)
	fmt.Print(strings.Join(strkeys, ","))
}
