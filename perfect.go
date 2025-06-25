package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var c float64
	fmt.Println("Enter the number")
	fmt.Scanln(&a)
	c = math.Sqrt(float64(a))

	if c == float64(int(c)) {
		fmt.Println("Perfect Square")
	} else {
		fmt.Println("Not perfect square")
	}
}
