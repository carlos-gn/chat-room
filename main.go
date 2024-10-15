package main

import (
	"chat_room/entities"
	"fmt"
)

func main() {
	u1 := entities.NewUser("Carlos")
	u2 := entities.NewUser("Trufa")

	r := entities.NewRoom("Room1")
	err := r.AddMessage("Good morning!", *u1)
	if err != nil {
		fmt.Println(err)
	}
	r.AddUser(u1)
	r.AddUser(u2)
	r.AddMessage("Guau", *u2)
	r.AddMessage("Guau", *u2)
	fmt.Printf("Latest messages: %+v  \n", r.GetLatestMessages(5))
	fmt.Printf("Room info: %+v", r.GetRoomInfo())
}
