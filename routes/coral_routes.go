package routes

import (
	"github.com/d11m08y03/digicup/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCoralRoutes(router *gin.Engine) {
	coralRoutes := router.Group("/coral")

	coralRoutes.GET("/about", func(ctx *gin.Context) {
		ctx.File("templates/about.html")
	})

	coralRoutes.GET("/contact", func(ctx *gin.Context) {
		ctx.File("templates/contact.html")
	})

	coralRoutes.GET("/home", func(ctx *gin.Context) {
		ctx.File("templates/coralhome.html")
	})

	coralRoutes.GET("/login", func(ctx *gin.Context) {
		ctx.File("templates/loginpage.html")
	})

	coralRoutes.Use(middleware.JWTAuthMiddleware())
	coralRoutes.GET("/dashboard", func(ctx *gin.Context) {
		ctx.File("templates/dashboard.html")
	})
}
