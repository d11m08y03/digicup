package main

import (
	"github.com/d11m08y03/digicup/config"
	"github.com/d11m08y03/digicup/middleware"
	"github.com/d11m08y03/digicup/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	router := gin.Default()
  router.Use(middleware.CORSMiddleware())

  routes.RegisterStaticRoutes(router)
	routes.RegisterUserRoutes(router)
  routes.RegisterAIRoutes(router)
  routes.RegisterCoralRoutes(router)

	router.Run(":8080")
}

