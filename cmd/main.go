package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Ginger API.")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
