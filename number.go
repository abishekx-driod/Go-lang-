package main

import "fmt"

func main() {
	var b int
	fmt.Println("Enter the number")
	fmt.Scanln(&b)
	num := 1
	for i := 1; i <= b; i++ {
		for j := 1; j <= i; j++ {
			if i%2 == 0 {
				fmt.Printf("%d", num)
				num--
			} else {
				fmt.Printf("%d", num)
				num++
			}
		}
		if (i+1)%2 == 0 {
			num = (num) + i
		} else {
			num = num + (i + 1)
		}
		fmt.Printf("\n")
	}
}
