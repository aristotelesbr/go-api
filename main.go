package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aristotelesbr/go-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
