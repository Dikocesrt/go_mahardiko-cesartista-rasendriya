package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	resData, err := io.ReadAll(res.Body)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	defer res.Body.Close()

	var users []User

	err = json.Unmarshal(resData, &users)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	for i, user := range users {
		fmt.Printf("Data User Ke-%v\n", i)
		fmt.Printf("user Id = %v\n", user.UserId)
		fmt.Printf("id = %v\n", user.Id)
		fmt.Printf("title = %v\n", user.Title)
		fmt.Printf("isCompleted = %v\n", user.Completed)
		fmt.Println("===================================")
	}
}