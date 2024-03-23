package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

	jsonStr := []byte(`{"userId": 2, "id": 2, "title": "new todo", "completed": true}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
	defer res.Body.Close()

	fmt.Println("Response Status: ", res.Status)
}