package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blessedmadukoma/trackit-chima/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file:", err)
	}

	port := os.Getenv("PORT")

	router := routes.Handlers()
	fmt.Printf("Server starting on port %s!\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
