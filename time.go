package main

import (
	"fmt"
	"time"
)

func main() {
	// Assuming image.Created is 2 days ago
	imageCreated := time.Now().Add(-48 * time.Hour).Unix()
	fmt.Println(imageCreated)
	b := time.Unix(0, imageCreated*int64(time.Millisecond)).Format("2006-01-02T15:04:05") + "Z"
	a := time.Unix(imageCreated, 0).Format("2006-01-02T15:04:05") + "Z"

	fmt.Println(a)
	fmt.Println(b)
}
