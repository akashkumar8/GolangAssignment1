package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {

	u1 := User{1, "Akash Kumar", "employee"}

	json_data, err := json.Marshal(u1)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(string(json_data))

	users := []User{
		{Id: 2, Name: "Verma Prem", Occupation: "trainee"},
		{Id: 3, Name: "Aaiky jordy", Occupation: "Singer"},
		{Id: 4, Name: "jordan dango", Occupation: "programmer"},
	}

	json_data2, err := json.Marshal(users)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(string(json_data2))
}
