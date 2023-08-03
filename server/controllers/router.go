package controllers

import (
	"typing/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	api := router.Group("/api")
	api.POST("/register", Register)
	api.POST("/login", Login)

	api.Use(middlewares.AuthorizeJwt())
	{
		api.GET("/user", GetAuthUser)

		api.GET("/scores", GetScores)
		api.POST("/scores", PostScore)
	}
	return router
}
