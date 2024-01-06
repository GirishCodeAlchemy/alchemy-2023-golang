package lib

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func FindSubstring() {
	// Step1
	parentDn := "sys/rack-unit-1/bios/bdgep/rack-unit-1"
	substr := "rack-unit-1"

	var indices []int
	index := strings.Index(parentDn, substr)

	for index != -1 {
		indices = append(indices, index)
		tmp := parentDn[index+len(substr):]
		fmt.Println("Next substring", tmp)
		index = strings.Index(tmp, substr)
		if index != -1 {
			index += len(substr) + indices[len(indices)-1]
		}
	}
	for _, v := range indices {
		result := parentDn[v : v+len(substr)]
		fmt.Println(v, result)
		if result == substr {
			fmt.Println("found")
		}

	}

	// Step2
	re := regexp.MustCompile(substr)
	fmt.Println("-->", re.FindAllIndex([]byte(parentDn), -1))
	fmt.Println("-->", re.FindAllString(parentDn, -1))
	fmt.Println("--->", re.FindAllStringIndex(parentDn, -1))
	fmt.Println("--->", re.ReplaceAllString(parentDn, "ggg"))
	fmt.Println([]byte(parentDn))

	fmt.Println(indices)

	numbers := []int{0, 42, 10, 8}
	fmt.Println(slices.Contains(numbers, 10))
}
func SubstringFunction(data string) []string {
	var result []string

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			result = append(result, data[i:j])
		}
	}
	fmt.Println(result)
	FindSubstring()
	return result
}
