package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"go-restaurant-management/middlewares"
	"go-restaurant-management/routes"
)

func main() {
	fmt.Println("Hello World")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
