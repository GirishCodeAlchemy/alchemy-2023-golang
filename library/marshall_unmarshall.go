package library

import (
	"encoding/json"
	"fmt"
)

type PersonData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func init() {
	fmt.Println("\n******** marshall_unmarshall smaple code *********")
	// Create a Person struct
	person := PersonData{Name: "Alice", Age: 30}

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("Marshalled JSON:", string(jsonData))

	// Unmarshal the JSON back into a struct
	var newPerson PersonData
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Unmarshalled Person:", newPerson)
}
