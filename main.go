package main

import "fmt"

type BaseEvent struct {
	AggregateStream string
}

type UserCreated struct {
	UserName string
	Password string
	Email    string
	BaseEvent
}

func main() {
	uc := UserCreated{
		UserName: "John Doe",
		Password: "1234",
		Email:    "john.doe@gmail.com",
		BaseEvent: BaseEvent{
			AggregateStream: "user",
		},
	}
	fmt.Println(uc.AggregateStream)
}
