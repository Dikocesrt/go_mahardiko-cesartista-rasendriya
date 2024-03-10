package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"name"`
}

func main() {
	url := "https://fakestoreapi.com/users"

	userCh := make(chan []User)
	doneCh := make(chan bool)

	var wg sync.WaitGroup

	getUsers := func(url string) {
		defer wg.Done()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()

		var users []User
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		userCh <- users
	}

	go func() {
		wg.Add(1)
		go getUsers(url)
		wg.Wait()
		close(userCh)
	}()

	go func() {
		fmt.Println("Users:")
		fmt.Println("====")
		for users := range userCh {
			for _, user := range users {
				fmt.Printf("ID: %d\nUsername: %s\nEmail: %s\nFirst name: %s\nLast name: %s\n====\n", user.ID, user.Username, user.Email, user.Name.FirstName, user.Name.LastName)
			}
		}
		doneCh <- true
	}()

	<-doneCh
}
