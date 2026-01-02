package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Ginger API.")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	log.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
