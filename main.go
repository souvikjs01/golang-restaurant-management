package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
func main() {
	fmt.Println("hello")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	router.Run(":" + port)
}
