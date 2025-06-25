package main

import (
	"fmt"
)

func main() {
	var b, i int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &b)
	num := 1
	for i = 1; i <= (b + b); i++ {
		k := i
		if i > b {
			k = (b + b) - i
		} else {
			k = i
		}
		for j := 1; j <= k; j++ {
			fmt.Printf("%d", num)
			num++
		}
		fmt.Printf("\n")
	}
}
