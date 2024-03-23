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
	url := "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	resData, err := io.ReadAll(res.Body)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	var users User

	err = json.Unmarshal(resData, &users)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	fmt.Println("---RESPONSE---")
	fmt.Printf("user Id = %v\n", users.UserId)
	fmt.Printf("id = %v\n", users.Id)
	fmt.Printf("title = %v\n", users.Title)
	fmt.Printf("isCompleted = %v\n", users.Completed)
	fmt.Println("===================================")
}