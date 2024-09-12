package main

import (
	"fmt"
	"log"
	"os"
	"product-service/config"
	"product-service/internal/infrastructure/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	r := http.SetupRoute(db)

	port := os.Getenv("PRODUCT_SERVICE_PORT")

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
