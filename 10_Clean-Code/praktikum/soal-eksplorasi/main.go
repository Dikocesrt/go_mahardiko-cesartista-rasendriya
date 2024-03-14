package main

import (
	"github.com/labstack/echo/v4"
)

type Request struct {
	OperandA int
	OperandB int
}

func main() {
	e := echo.New()

	e.POST("/api/add", func(c echo.Context) error {
		req := new(Request)

		c.Bind(req)

		return c.JSON(200, echo.Map{
			"result": req.OperandA + req.OperandB,
		})
	})

	e.POST("/api/subtract", func(c echo.Context) error {
		req := new(Request)

		c.Bind(req)

		return c.JSON(200, echo.Map{
			"result": req.OperandA - req.OperandB,
		})
	})

	e.POST("/api/multiply", func(c echo.Context) error {
		req := new(Request)

		c.Bind(req)

		return c.JSON(200, echo.Map{
			"result": req.OperandA * req.OperandB,
		})
	})

	e.POST("/api/divide", func(c echo.Context) error {
		req := new(Request)

		c.Bind(req)

		// Handling division by zero
		if req.OperandB == 0 {
			return c.JSON(400, echo.Map{"error": "division by zero"})
		}

		return c.JSON(200, echo.Map{
			"result": req.OperandA / req.OperandB,
		})
	})

	e.Start(":1323")
}
