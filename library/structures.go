package library

import (
	"fmt"
	"unsafe"
)

// method struct

type Employee struct {
	Name string
	Age  int
}

func (e Employee) Display() {
	e.Age++
	fmt.Println("Hello", e.Name, e.Age)
}

// pointers methods

type EmployeePointers struct {
	Name string
	Age  int
}

func (e *EmployeePointers) Display() {
	e.Age++
	fmt.Println("Hello", e.Name, e.Age)
}

// type composition sturct
type EmployeeComposition struct {
	emp  Employee
	Name string
}

func Emptystruct() {
	mapobj := make(map[string]struct{})
	for _, value := range []string{"interviewbit", "golang", "questions"} {
		mapobj[value] = struct{}{}
	}
	fmt.Println(mapobj)

}

func StructureFunction() {

	ad := struct{}{} //empty struct
	fmt.Println(ad)
	Emptystruct()
	println("---empty struct", unsafe.Sizeof(ad))
	type person struct {
		name string
		age  int
	}
	emp := new(person)
	fmt.Println(emp)
	emp.name = "Raj"
	emp.age = 23
	fmt.Println(emp)

	emp1 := person{}
	emp1.name = "Raj1"
	emp1.age = 24
	fmt.Println(emp1)

	emp2 := person{name: "giri", age: 25}
	fmt.Println(emp2)

	emp3 := person{age: 26}
	fmt.Println(emp3)

	emp4 := Employee{Name: "giri", Age: 25}
	emp4.Display()
	emp4.Display()

}
