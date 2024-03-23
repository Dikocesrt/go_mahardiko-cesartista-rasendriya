package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
	defer res.Body.Close()

	fmt.Println("Response Status: ", res.Status)
}