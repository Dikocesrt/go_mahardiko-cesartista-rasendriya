package main

import (
	"bytes"
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

	jsonStr := []byte(`{"client_id": 2, "client_key": "CLIENT03", "client_secret": "new secret", "status": true}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

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