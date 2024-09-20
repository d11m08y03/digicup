package routes

import "github.com/gin-gonic/gin"

func RegisterStaticRoutes(router *gin.Engine) {
	router.Static("/static", "./static")
}
