package configs

import (
	"clean/repositories/mysql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfigMySQL() mysql.Config {
	return mysql.Config{
		DBName: os.Getenv("DBName"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBUnix: os.Getenv("DBUnix"),
	}
}