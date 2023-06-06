package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controllers "gin/controllers"
	middleware "gin/middleware"
	routes "gin/routes"
)

func main() {
	controllers.Connect()
	router := gin.Default()

	config := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})
	router.Use(config)

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthMiddleware)
	routes.PrivateRoutes(private)

	router.Run("localhost:8080")
}
