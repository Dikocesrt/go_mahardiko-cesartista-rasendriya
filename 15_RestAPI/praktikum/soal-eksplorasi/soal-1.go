package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Fruit struct {
    Name       string `json:"name"`
    ID         int    `json:"id"`
    Family     string `json:"family"`
    Order      string `json:"order"`
    Genus      string `json:"genus"`
    Nutritions struct {
        Calories     int `json:"calories"`
        Fat          float64 `json:"fat"`
        Sugar        float64 `json:"sugar"`
        Carbohydrates float64 `json:"carbohydrates"`
        Protein      float64 `json:"protein"`
    } `json:"nutritions"`
}

func main() {
	fmt.Print("enter fruit name : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var fruit = scanner.Text()
	var url = "https://www.fruityvice.com/api/fruit/" + fruit

	res, err := http.Get(url)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	resData, err := io.ReadAll(res.Body)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	defer res.Body.Close()

	var fruits Fruit

	err = json.Unmarshal(resData, &fruits)

	if err != nil {
        log.Fatalf("Error fetching data: %v", err)
    }

	fmt.Println("== FRUIT DATA ==")
	fmt.Printf("Fruit name: %v\n", fruits.Name)
	fmt.Printf("Fruit family: %v\n", fruits.Family)
	fmt.Printf("Calories: %v\n", fruits.Nutritions.Calories)
	fmt.Printf("Fat: %v\n", fruits.Nutritions.Fat)
	fmt.Printf("Sugar: %v\n", fruits.Nutritions.Sugar)
	fmt.Printf("Carbohydrates: %v\n", fruits.Nutritions.Carbohydrates)
	fmt.Printf("Protein: %v\n", fruits.Nutritions.Protein)
}