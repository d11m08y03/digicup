package routes

import (
	"github.com/d11m08y03/digicup/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAIRoutes(router *gin.Engine) {
	aiRoutes := router.Group("/ai")

	aiRoutes.POST("/upload", controllers.UploadImage)
}
