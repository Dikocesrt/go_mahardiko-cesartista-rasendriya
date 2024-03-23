package main

import (
	"fmt"
	"net/http"
)

type client struct {
	Client_id     int    `json:"client_id"`
	Client_key    string `json:"client_key"`
	Client_secret string `json:"client_secret"`
	Status        bool   `json:"status"`
}

func main() {
	url := "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/client/2"

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