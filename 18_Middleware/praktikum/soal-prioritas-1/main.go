package main

import (
	"middleware-prioritas-1/config"
	"middleware-prioritas-1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}