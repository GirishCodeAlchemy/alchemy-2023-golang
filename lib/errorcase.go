package lib

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: negative number passed to sort")
		// return 0, fmt.Errorf("Math: negative number passed to sort: %v", x)
	}
	return math.Sqrt(x), nil
}

func ErrorCase() {
	res, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
