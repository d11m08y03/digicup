package routes

import "github.com/gin-gonic/gin"

func RegisterStaticRoutes(router *gin.Engine) {
	router.GET("/about", func(ctx *gin.Context) {
		ctx.File("templates/about.html")
	})

	router.GET("/contact", func(ctx *gin.Context) {
		ctx.File("templates/contact.html")
	})

	router.GET("/coralHome", func(ctx *gin.Context) {
		ctx.File("templates/coralhome.html")
	})

	router.GET("/dashboard", func(ctx *gin.Context) {
		ctx.File("templates/dashboard.html")
	})

	router.GET("/login", func(ctx *gin.Context) {
		ctx.File("templates/loginpage.html")
	})

	router.Static("/static", "./static")
}
