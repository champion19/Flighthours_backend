package main

import (
	"log"
	"os"

	"github.com/champion19/Flighthours_backend/cmd/api"
	"github.com/gin-gonic/gin"
)

func main() {
	err := os.Setenv("PORT", "8081")
	app := gin.Default()
	api.Boostrap(app)
	if err := app.Run(); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalf("failed to set env: %v", err)
	}
}
