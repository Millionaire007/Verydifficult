package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelsequeira07/ecommerce-mic/controllers"
	"github.com/michaelsequeira07/ecommerce-mic/database"
	"github.com/michaelsequeira07/ecommerce-mic/middleware"
	"github.com/michaelsequeira07/ecommerce-mic/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApllication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
