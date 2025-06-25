package main

import (
	"fmt"
	"sort"
)

func main() {
	var name string = "45"
	//var name2 string = "456"
	fmt.Printf("Name is %X", name)
	fmt.Println("Hello Abhi")
	fmt.Println("Abiii")
	fmt.Println("The Bill Amount is")
	names := []string{"Abhi", "vain", "bha"}
	fmt.Println(names)
	names[2] = "vya"
	fmt.Println(names)
	names = append(names, "ss")
	fmt.Println(names[1:2], len(names))
	ages := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(sort.SearchInts(ages, 23))

}
