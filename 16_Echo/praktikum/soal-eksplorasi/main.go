package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Fruit struct {
	ID         int    `json:"id"`
    Name       string `json:"name"`
	Price int `json:"price"`
    Nutritions struct {
        Calories     int `json:"calories"`
        Fat          float64 `json:"fat"`
        Sugar        float64 `json:"sugar"`
        Carbohydrates float64 `json:"carbohydrates"`
        Protein      float64 `json:"protein"`
    } `json:"nutritions"`
}

type BaseResponseError struct {
	Message   string `json:"message"`
}

type BaseResponseSuccess struct {
	Data    any    `json:"data"`
}

type App struct {
	DB *sql.DB
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_fruits_alterra")

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}	

	defer db.Close()

	fmt.Println("Connected to database")

	app := &App{DB: db}

	e := echo.New()
	e.GET("/api/v1/fruits", app.getAllFruitsHandler)
	e.GET("/api/v1/fruits/:id", app.getFruitByIdHandler)
	e.POST("/api/v1/fruits", app.addFruitHandler)
	e.DELETE("/api/v1/fruits/:id", app.deleteFruitByIdHandler)
	e.Start(":8080")
}

func (app *App) deleteFruitByIdHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	rows, err := app.DB.Query("DELETE from fruits WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted"})
}

func (app *App) getFruitByIdHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	rows, err := app.DB.Query("SELECT * FROM fruits WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()

	var fruits []Fruit

	for rows.Next(){
		var fruit Fruit

		if err := rows.Scan(&fruit.ID, &fruit.Name, &fruit.Price, &fruit.Nutritions.Calories, &fruit.Nutritions.Fat, &fruit.Nutritions.Sugar, &fruit.Nutritions.Carbohydrates, &fruit.Nutritions.Protein); err != nil {
			fmt.Println(err)
			return err
		}

		fruits = append(fruits, fruit)
	}

	baseResponseSuccess := BaseResponseSuccess {
		Data: fruits,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func (app *App) addFruitHandler(c echo.Context) error {
	var fruit Fruit
	c.Bind(&fruit)

	var url = "https://www.fruityvice.com/api/fruit/" + fruit.Name

	res, _ := http.Get(url)
	resData, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(resData, &fruit)

	id := int(uuid.New().ID())

	fruit.ID = id

	_, err := app.DB.Exec("INSERT INTO fruits (id, nama, harga, kalori, lemak, gula, karbohidrat, protein) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", fruit.ID, fruit.Name, fruit.Price, fruit.Nutritions.Calories, fruit.Nutritions.Fat, fruit.Nutritions.Sugar, fruit.Nutritions.Carbohydrates, fruit.Nutritions.Protein)

	if err != nil {
		fmt.Println(err)
		return err
	}

	baseResponseSuccess := BaseResponseSuccess {
		Data: fruit,
	}

	return c.JSON(http.StatusCreated, baseResponseSuccess)
}

func (app *App) getAllFruitsHandler(c echo.Context) error {
	rows, err := app.DB.Query("SELECT * FROM fruits")

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()

	var fruits []Fruit

	for rows.Next(){
		var fruit Fruit

		if err := rows.Scan(&fruit.ID, &fruit.Name, &fruit.Price, &fruit.Nutritions.Calories, &fruit.Nutritions.Fat, &fruit.Nutritions.Sugar, &fruit.Nutritions.Carbohydrates, &fruit.Nutritions.Protein); err != nil {
			fmt.Println(err)
			return err
		}

		fruits = append(fruits, fruit)
	}

	baseResponseSuccess := BaseResponseSuccess {
		Data: fruits,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}