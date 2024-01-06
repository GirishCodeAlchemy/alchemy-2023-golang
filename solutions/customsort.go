// sort a slice of custom structs with the help of an example
package solutions

import (
	"fmt"
	"sort"
)

type Humans struct {
	name string
	age  int
}

// Sort by number
type ByAge []Humans

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort by string
type ByName []Humans

func (n ByName) Len() int           { return len(n) }
func (n ByName) Less(i, j int) bool { return n[i].name < n[j].name }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// sort the value
type KeyValue struct {
	Key   string
	Value int
}

// sort slice
type seq []int

func (s seq) Len() int           { return len(s) }
func (s seq) Less(i, j int) bool { return s[i] < s[j] }
func (s seq) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func CustomSort() {
	audience := []Humans{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
		{"Cecil", 15},
	}
	fmt.Println(audience)
	sort.Sort(ByAge(audience))
	fmt.Println(audience)
	sort.Sort(ByName(audience))
	fmt.Println(audience)
	// Stable slice
	CustomComparator()

	// sort mechanism
	b := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(b)
	fmt.Println(b)

	intseq := []int{20, 10, 30, 40, 15, 14, 16, 12}
	fmt.Printf("\nUnsorted List: %v\n", intseq)
	sort.Sort(seq(intseq))
	fmt.Printf("sorted List %v\n", intseq)

	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	var keyValueSlice []KeyValue
	for k, v := range m {
		keyValueSlice = append(keyValueSlice, KeyValue{k, v})
	}
	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Value < keyValueSlice[j].Value
	})
	fmt.Println(keyValueSlice)
	for _, v := range keyValueSlice {
		fmt.Println(v.Key, v.Value)
	}

}

// sort a by stable

func CustomComparator() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)
}
