package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/maudev1/meu-primeiro-crud/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {

		log.Fatalf("Error loading .env file: %v", err)

		os.Exit(1)
	}

	router := gin.Default()
	// web := r.Group("/")
	routes.InitRoutes(&router.RouterGroup)
	// if(router.Run(":8080"): err != nil)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
