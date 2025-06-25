package main

import "fmt"

func main() {
	var b int
	fmt.Println("Enter the number")
	fmt.Scanln(&b)
	for i := 1; i <= b; i++ {
		for j := 1; j <= (b - i); j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= (i+i)-1; k++ {
			if k%2 == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", i)
			}
		}
		fmt.Printf("\n")
	}
}
