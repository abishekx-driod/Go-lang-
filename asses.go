package main

import (
	"fmt"
)

func SquareEvens(nums []int) []int {
	ress := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			ress = append(ress, nums[i]*nums[i])
		}
	}
	return ress
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	result := SquareEvens(input)
	fmt.Println("Squared Evens:", result)
}
