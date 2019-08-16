package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// a := []int{5, 4, 3, 2, 1}
	// a = append(a, 13)
	// fmt.Println(a)

	// vertices := make(map[string]int)

	// vertices["triangle"] = 2
	// vertices["square"] = 3
	// vertices["dodecagon"] = 12

	// delete(vertices, "square")

	// fmt.Println(vertices)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// m := make(map[string]string)
	// m["a"] = "alpha"
	// m["b"] = "beta"

	// for key, value := range m {
	// 	fmt.Println("key", key, "value:", value)
	// }
	// 	result := sum(2, 3)
	// 	fmt.Println(result)
	// }
	// func sum(x int, y int) int {
	// 	return x + y

	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefinded for negative numbers")
	}
	return math.Sqrt(x), nil

}
