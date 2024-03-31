package main

import (
	"middleware-prioritas-2/config"
	"middleware-prioritas-2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}