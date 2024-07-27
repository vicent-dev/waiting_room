package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	env := os.Getenv("WAITING_ROOM_ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(".env." + env)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
