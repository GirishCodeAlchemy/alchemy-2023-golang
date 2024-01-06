package solutions

import "fmt"

// Person is a simple struct representing a person
type Person struct {
	FirstName, LastName string
}

// Introduction returns a greeting using a value receiver
func (p Person) Introduction() string {
	return fmt.Sprintf("Hello, I'm %s %s", p.FirstName, p.LastName)
}

// ChangeLastName updates the last name using a pointer receiver
func (p *Person) ChangeLastName(newLastName string) {
	p.LastName = newLastName
}

func recieverFunction() {
	person := Person{FirstName: "John", LastName: "Doe"}

	// Using a value receiver
	greeting := person.Introduction()
	fmt.Println(greeting)

	// Using a pointer receiver to modify the object
	person.ChangeLastName("Smith")
	fmt.Printf("Updated Last Name: %s\n", person.LastName)
}

func init() {

	recieverFunction()

}
