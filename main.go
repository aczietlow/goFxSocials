package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")

	// debug BS
	fmt.Println(token)
}
