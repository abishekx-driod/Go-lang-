package main

import "fmt"

func main() {
	var sr string
	fmt.Scanln(&sr)
	ma := make(map[rune]int)
	for _, r := range sr {
		ma[r]++
	}
	co := -1
	for i, r := range sr {
		if ma[r] == 1 {
			co = i
			break
		}
	}
	fmt.Println((co))
}
