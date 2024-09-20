package main

import (
	"github.com/d11m08y03/digicup/config"
	"github.com/d11m08y03/digicup/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	router := gin.Default()
	routes.RegisterUserRoutes(router)
  routes.RegisterStaticRoutes(router)

	router.Run(":8080")
}
