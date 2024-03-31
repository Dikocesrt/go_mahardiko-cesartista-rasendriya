package main

import (
	"middleware-eksplorasi/config"
	"middleware-eksplorasi/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}