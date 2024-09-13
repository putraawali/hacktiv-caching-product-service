package main

import (
	"fmt"
	"log"
	"os"
	"product-service/internal/infrastructure/http"
	"time"

	"github.com/gin-contrib/cache/persistence"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env file")
	}

	// Inisialisasi cache store dengan expiration time
	store := persistence.NewInMemoryStore(time.Second * 20)

	r := http.SetupRoute(store)

	port := os.Getenv("PRODUCT_SERVICE_PORT")

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
