package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	id := getUserById("Abi")
	fmt.Println(id)
	go getUserByChats()
	//log.Println(chats)
	go getUserByFriends()
	//log.Println(friends)
	log.Println(time.Since(time.Now()))
	//time.Since(start)
}
func getUserByChats() []string {
	time.Sleep(time.Second * 4)
	return []string{
		"Abi",
		"Siva",
		"Sing",
	}
}
func getUserByFriends() []string {
	time.Sleep(time.Second * 6)
	return []string{
		"Abi",
	}
}
func getUserById(na string) string {
	return fmt.Sprintf("%s-2", na)
}
