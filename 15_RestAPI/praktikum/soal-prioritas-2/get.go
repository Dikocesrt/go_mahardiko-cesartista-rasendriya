package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Client_id     int    `json:"client_id"`
	Client_key    string `json:"client_key"`
	Client_secret string `json:"client_secret"`
	Status        bool   `json:"status"`
}

func main() {
	url := "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client/"

	req, err := http.NewRequest("GET", url, nil)
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

	var clients []Client

	err = json.NewDecoder(res.Body).Decode(&clients)

	for i, client := range clients {
		fmt.Printf("Data User Ke-%v\n", i)
		fmt.Printf("Client Id = %v\n", client.Client_id)
		fmt.Printf("Client Key = %v\n", client.Client_key)
		fmt.Printf("Client Secret = %v\n", client.Client_secret)
		fmt.Printf("Status = %v\n", client.Status)
	}
}