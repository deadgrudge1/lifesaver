package main

import (
	"lifesaver/internal/lifesaver/rest"
	"github.com/gin-gonic/gin"
)

func initRouter() {
    router := gin.Default()

	gin.Logger()
	router.Use(gin.Logger())

	//User Endpoints
    router.GET("/user/:userId", rest.GetUser)
	router.POST("/user", rest.SaveUser)
	router.PUT("/user/:userId", rest.UpdateUser)
	router.DELETE("/user/:userId", rest.DeleteUser)

    router.Run("localhost:8081")
}