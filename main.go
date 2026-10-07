package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Run(":" + port)

}
