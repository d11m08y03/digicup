package routes

import (
	"github.com/d11m08y03/digicup/controllers"
	"github.com/d11m08y03/digicup/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	userRoutes.POST("/register", controllers.Register)
	userRoutes.POST("/login", controllers.Login)

	userRoutes.Use(middleware.JWTAuthMiddleware())
	userRoutes.GET("/getUsers", controllers.GetUsers)
}
