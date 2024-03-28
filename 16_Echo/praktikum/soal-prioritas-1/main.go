package main

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Foods struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
}

type BaseResponseError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Status    bool   `json:"status"`
}

type BaseResponseSuccess struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}

func main() {
	e := echo.New()
	e.GET("/api/v1/foods", getFoodsHandler)
	e.GET("/api/v1/foods/:id", getDetailedFoodHandler)
	e.POST("/api/v1/foods", addFoodHandler)
	e.PUT("/api/v1/foods/:id", updateFoodHandler)
	e.DELETE("/api/v1/foods/:id", deleteFoodHandler)
	e.Start(":8080")
}

func deleteFoodHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	errResponse := BaseResponseError{}
	if err != nil {
		errResponse.ErrorCode = 4031
		errResponse.Message = "Error Brow"
		errResponse.Status = false
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	foods := generateFoods()

	indexToDelete := -1
	for i, food := range foods {
		if food.Id == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		foods = append(foods[:indexToDelete], foods[indexToDelete+1:]...)
	} else {
		errResponse.ErrorCode = 4031
		errResponse.Message = "Id tidak ditemukan"
		errResponse.Status = false
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	baseResponseSuccess := BaseResponseSuccess {
		Message: "Success",
		Status: true,
		Data: foods,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func updateFoodHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var foodBody Foods
	c.Bind(&foodBody)

	foods := []Foods{
		{1, "Nasi Goreng", 15000, "Nasi goreng dengan bumbu rempah khas Indonesia"},
		{2, "Soto Ayam", 12000, "Soto ayam dengan kuah kaldu ayam yang gurih"},
		{int(uuid.New().ID()), "Mie Ayam", 13000, "Mie ayam dengan potongan ayam dan kuah kaldu"},
		{int(uuid.New().ID()), "Gado-Gado", 16000, "Sayuran segar dengan bumbu kacang khas Indonesia"},
		{int(uuid.New().ID()), "Bakso", 10000, "Bakso dengan kuah kaldu sapi dan tambahan mie"},
	}
	
	for i, food := range foods {
		if id == food.Id {
			foods[i].Name = foodBody.Name
			foods[i].Description = foodBody.Description
			foods[i].Price = foodBody.Price
		}
	}

	baseResponseSuccess := BaseResponseSuccess {
		Message: "Success",
		Status: true,
		Data: foods,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func addFoodHandler(c echo.Context) error {
	food := Foods{}
	c.Bind(&food)

	food.Id = int(uuid.New().ID())

	baseResponseSuccess := BaseResponseSuccess {
		Message: "Success",
		Status: true,
		Data: food,
	}

	return c.JSON(http.StatusCreated, baseResponseSuccess)
}

func getDetailedFoodHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	errResponse := BaseResponseError{}

	if err != nil {
		errResponse.ErrorCode = 4031
		errResponse.Message = "Error Brow"
		errResponse.Status = false
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	foods := generateFoods()
	var resultData Foods

	for _, food := range foods {
		if food.Id == idInt {
			resultData = food
		}
	}

	baseResponseSuccess := BaseResponseSuccess {
		Message: "Success",
		Status: true,
		Data: resultData,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func getFoodsHandler(c echo.Context) error {
	foods := generateFoods()

	baseResponseSuccess := BaseResponseSuccess {
		Message: "Success",
		Status: true,
		Data: foods,
	}

	return c.JSON(http.StatusOK, baseResponseSuccess)
}

func generateFoods() []Foods {
	foods := []Foods{
		{1, "Nasi Goreng", 15000, "Nasi goreng dengan bumbu rempah khas Indonesia"},
		{int(uuid.New().ID()), "Soto Ayam", 12000, "Soto ayam dengan kuah kaldu ayam yang gurih"},
		{int(uuid.New().ID()), "Mie Ayam", 13000, "Mie ayam dengan potongan ayam dan kuah kaldu"},
		{int(uuid.New().ID()), "Gado-Gado", 16000, "Sayuran segar dengan bumbu kacang khas Indonesia"},
		{int(uuid.New().ID()), "Bakso", 10000, "Bakso dengan kuah kaldu sapi dan tambahan mie"},
	}

	return foods
}